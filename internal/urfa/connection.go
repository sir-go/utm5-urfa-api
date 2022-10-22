package urfa

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/binary"
	"net"
	"time"

	"github.com/op/go-logging"
	"github.com/pkg/errors"
	"github.com/zhuangsirui/binpacker"
)

var log = logging.MustGetLogger("urfa-go")

type Connection struct {
	conn     net.Conn
	unpacker *binpacker.Unpacker
	packer   *binpacker.Packer

	inPkt      *Packet
	curAttrIdx int

	outPkt *Packet
}

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

func (c *Connection) Disconnect() error {
	return c.conn.Close()
}

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

	if err = c.receive([]uint8{PacketTypeSessionInit}); err != nil {
		return err
	}

	sessionKeyAttr := c.inPkt.GetAttr(AttrTypeKey)
	if sessionKeyAttr == nil {
		return errors.New("SessionInit key is empty")
	}
	authHash := md5.New()
	authHash.Write(sessionKeyAttr.Data)
	authHash.Write([]byte(password))

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
	c.Send()

	if err = c.receive([]uint8{PacketTypeAccessAccept, PacketTypeAccessReject}); err != nil {
		return err
	}

	if c.inPkt.Header.Type == PacketTypeAccessReject {
		return errors.Errorf("Login request rejected")
	}

	return c.wrapSSL(cert)
}

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

func (c *Connection) Send() {
	defer func() {
		if err := c.conn.SetWriteDeadline(time.Time{}); err != nil {
			log.Error(err)
		}
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

func (c *Connection) receive(expectedPacketTypes []uint8) (err error) {
	defer func() {
		if err := c.conn.SetReadDeadline(time.Time{}); err != nil {
			log.Error(err)
		}
	}()

	c.inPkt = new(Packet)
	c.unpacker.FetchUint8(&c.inPkt.Header.Type).FetchUint8(&c.inPkt.Header.Ver).FetchUint16(&c.inPkt.Header.Len)
	if err = c.unpacker.Error(); err != nil {
		return
	}

	if c.inPkt.Header.Len < PacketHeaderSize+AttrHeaderSize {
		log.Debug("packet has no attributes")
		return
	}

	if !packetTypeIsExpected(c.inPkt.Header.Type, expectedPacketTypes) {
		return errors.Errorf("unexpected packet type: %d", c.inPkt.Header.Type)
	}

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

	// log.Debugf("< %v", pkt)
	return
}

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

func (c *Connection) getRaw() []byte {
	if c.inPkt == nil || c.curAttrIdx+2 > len(c.inPkt.Attributes) {
		c.fetchPacket()
		return c.getRaw()
	}
	c.curAttrIdx++

	return c.inPkt.Attributes[c.curAttrIdx].Data
}

func (c *Connection) putData(data []byte) {
	if c.outPkt == nil {
		c.outPkt = NewPacket(PacketTypeSessionData)
	}
	c.outPkt.AddAttr(AttrTypeData, data)
}

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
