package urfa

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/binary"
	"net"
	"time"

	"github.com/pkg/errors"
	"github.com/zhuangsirui/binpacker"
)

// Connection holds a TCP connection to the billing server and processes packets
type Connection struct {
	conn     net.Conn            // TCP connection to the URFA server
	unpacker *binpacker.Unpacker // read a data from the socket and unpacks it
	packer   *binpacker.Packer   // packs and write to the socket a data

	// input packet pointer
	inPkt *Packet

	// packet attribute pointer
	curAttrIdx int

	// output packet pointer
	outPkt *Packet
}

// Connect makes a Connection to the URFA server and initializes packer and unpacker
func Connect(addr string, username string, password string, cert string) (uConn *Connection, err error) {
	uConn = new(Connection)
	uConn.conn, err = net.DialTimeout("tcp", addr, time.Minute*1)
	if err != nil {
		return nil, err
	}

	uConn.unpacker = binpacker.NewUnpacker(binary.BigEndian, uConn.conn)
	uConn.packer = binpacker.NewPacker(binary.BigEndian, uConn.conn)

	return uConn, uConn.login(username, password, cert)
}

// Disconnect closes a Connection
func (c *Connection) Disconnect() error {
	return c.conn.Close()
}

// Login sends an Auth packet to the server and wrap the connection with SSL if success
func (c *Connection) login(username string, password string, cert string) (err error) {
	defer func() {
		var isError bool
		if r := recover(); r != nil {
			err, isError = r.(error)
			if !isError {
				err = errors.New("unknown login error")
			}
		}
	}()

	// read the Hello packet with a Session init key from the server
	if err = c.receive([]uint8{PacketTypeSessionInit}); err != nil {
		return err
	}

	sessionKeyAttr := c.inPkt.GetAttr(AttrTypeKey)
	if sessionKeyAttr == nil {
		return errors.New("SessionInit key is empty")
	}

	// encrypt the password (yes, the old and weak md5 ¯\_(ツ)_/¯
	authHash := md5.New()
	authHash.Write(sessionKeyAttr.Data)
	authHash.Write([]byte(password))

	// make a new packet
	c.outPkt = NewPacket(PacketTypeAccessRequest)

	loginTypeVal := make([]byte, 4)
	binary.BigEndian.PutUint32(loginTypeVal, LoginTypeSystem)
	c.outPkt.AddAttr(AttrTypeLoginType, loginTypeVal)
	c.outPkt.AddAttr(AttrTypeLogin, []byte(username))
	c.outPkt.AddAttr(AttrTypeCHAPChallenge, sessionKeyAttr.Data)
	c.outPkt.AddAttr(AttrTypeCHAPResponse, authHash.Sum(nil))

	sslReqTypeVal := make([]byte, 4)
	binary.BigEndian.PutUint32(sslReqTypeVal, SslReqTypeTLS12)
	c.outPkt.AddAttr(AttrTypeSSLRequest, sslReqTypeVal)

	// send auth packet
	c.Send()

	// receive the answer
	if err = c.receive([]uint8{PacketTypeAccessAccept, PacketTypeAccessReject}); err != nil {
		return err
	}

	// if rejected - break
	if c.inPkt.Header.Type == PacketTypeAccessReject {
		return errors.Errorf("Login request rejected")
	}

	// wrap the connection
	return c.wrapSSL(cert)
}

// wrapSSL reads a cert file, wraps the connection with SSL and reinitialize packer and unpacker
func (c *Connection) wrapSSL(certPath string) (err error) {
	cert, err := tls.LoadX509KeyPair(certPath, certPath)
	if err != nil {
		return err
	}

	tlsConf := tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	tlsConn := tls.Client(c.conn, &tlsConf)
	if err = tlsConn.Handshake(); err != nil {
		_ = tlsConn.Close()
		return err
	}

	c.unpacker = binpacker.NewUnpacker(binary.BigEndian, tlsConn)
	c.packer = binpacker.NewPacker(binary.BigEndian, tlsConn)

	return
}

// Send writes an output packet data to the packer and clears an input packet to receive the answer
func (c *Connection) Send() {
	defer func() {
		_ = c.conn.SetWriteDeadline(time.Time{})
	}()

	buf := bytes.NewBuffer([]byte{})
	packer := binpacker.NewPacker(binary.BigEndian, buf)

	// log.Debugf("> %v", pkt)
	if c.outPkt == nil {
		c.outPkt = NewPacket(PacketTypeSessionData)
	}
	if c.outPkt.Header.Len == 0 {
		c.outPkt.Header.Len = PacketHeaderSize
	}
	packer.PushUint8(c.outPkt.Header.Type).PushUint8(PacketVersion).PushUint16(c.outPkt.Header.Len)

	for _, attr := range c.outPkt.Attributes {
		packer.PushUint8(
			attr.Header.Type).PushUint8(
			attr.Header.Ver).PushUint16(
			attr.Header.Len).PushBytes(
			attr.Data)
	}

	if err := packer.Error(); err != nil {
		panic(err)
	}

	c.packer.PushBytes(buf.Bytes())
	if err := packer.Error(); err != nil {
		panic(err)
	}

	c.outPkt = nil
	return
}

// receive reads an input packet data (all of attributes) from the unpacker
func (c *Connection) receive(expectedPacketTypes []uint8) (err error) {
	defer func() {
		_ = c.conn.SetReadDeadline(time.Time{})
	}()

	// prepare the input packet
	c.inPkt = new(Packet)
	c.unpacker.FetchUint8(&c.inPkt.Header.Type).FetchUint8(&c.inPkt.Header.Ver).FetchUint16(&c.inPkt.Header.Len)
	if err = c.unpacker.Error(); err != nil {
		return
	}

	if c.inPkt.Header.Len < PacketHeaderSize+AttrHeaderSize {
		return
	}

	if !packetTypeIsExpected(c.inPkt.Header.Type, expectedPacketTypes) {
		return errors.Errorf("unexpected packet type: %d", c.inPkt.Header.Type)
	}

	// sequentially read all attributes
	leftBytes := c.inPkt.Header.Len - PacketHeaderSize
	for leftBytes >= AttrHeaderSize {

		attr := new(Attribute)

		if err = c.conn.SetReadDeadline(time.Now().Add(time.Second * 10)); err != nil {
			return
		}

		c.unpacker.FetchUint8(&attr.Header.Type).FetchUint8(&attr.Header.Ver).FetchUint16(&attr.Header.Len)
		if err = c.unpacker.Error(); err != nil {
			return
		}

		if err = c.conn.SetReadDeadline(time.Now().Add(time.Second * 10)); err != nil {
			return
		}

		c.unpacker.FetchBytes(uint64(attr.Header.Len-AttrHeaderSize), &attr.Data)
		if err = c.unpacker.Error(); err != nil {
			return
		}

		leftBytes -= attr.Header.Len
		c.inPkt.Attributes = append(c.inPkt.Attributes, *attr)
	}
	return
}

// packetTypeIsExpected checks if gotten packet type is expected
func packetTypeIsExpected(pType uint8, expected []uint8) bool {
	if expected == nil {
		return true
	}
	for _, expType := range expected {
		if pType == expType {
			return true
		}
	}
	return false
}

// Call initiates an RPC session
func (c *Connection) Call(fnCode int) {
	var err error

	c.outPkt = NewPacket(PacketTypeSessionCall)

	codeBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(codeBytes, uint32(fnCode))
	c.outPkt.AddAttr(AttrTypeCall, codeBytes)
	c.Send()

	if err = c.receive([]uint8{PacketTypeSessionData}); err != nil {
		panic(err)
	}

	if err = c.checkAnswer(); err != nil {
		panic(err)
	}

	c.inPkt = nil
}

// fetchPacket gets an input packet and checks if it's a Termination packet
func (c *Connection) fetchPacket() {
	var err error

	if err = c.receive([]uint8{PacketTypeSessionData}); err != nil {
		panic(err)
	}

	if err = c.checkAnswer(); err != nil {
		panic(err)
	}

	c.curAttrIdx = -1
}

// getRaw recursively gets all the data from the packet
func (c *Connection) getRaw() []byte {
	if c.inPkt == nil || c.curAttrIdx+2 > len(c.inPkt.Attributes) {
		c.fetchPacket()
		return c.getRaw()
	}
	c.curAttrIdx++

	return c.inPkt.Attributes[c.curAttrIdx].Data
}

// putData makes a SessionData packet and puts given data to the Data attribute of it
func (c *Connection) putData(data []byte) {
	if c.outPkt == nil {
		c.outPkt = NewPacket(PacketTypeSessionData)
	}
	c.outPkt.AddAttr(AttrTypeData, data)
}

// checkAnswer checks the input packet type and makes an error if it Termination
func (c *Connection) checkAnswer() error {
	if termAttr := c.inPkt.GetAttr(AttrTypeTermination); termAttr != nil {
		val := binary.BigEndian.Uint32(termAttr.Data)
		switch val {
		case TermValueAccessRejected:
			return errors.New("access rejected")
		case TermValueIllegalCode:
			return errors.New("illegal function code")
		case TermValueForbidden:
			return errors.New("forbidden")
		case TermValueDone:
			return errors.New("got terminated response")
		default:
			return errors.Errorf("unknown termination attribute value: %d", val)
		}
	}
	return nil
}
