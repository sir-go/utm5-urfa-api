package urfa

import (
	"encoding/base64"
	"encoding/binary"
	"math"
	"net"
)

// Connection low-level IO operations for several scalar types

// ------------------------------------------  get

func (c *Connection) GetS() string {
	return string(c.getRaw())
}

func (c *Connection) GetI() int {
	return int(int32(binary.BigEndian.Uint32(c.getRaw())))
}

func (c *Connection) GetL() int64 {
	return int64(binary.BigEndian.Uint64(c.getRaw()))
}

func (c *Connection) GetD() float64 {
	return math.Float64frombits(binary.BigEndian.Uint64(c.getRaw()))
}

func (c *Connection) GetA() net.IP {
	rawBytes := c.getRaw()
	if len(rawBytes) < 5 {
		return nil
	}
	return rawBytes[1:]
}

func (c *Connection) GetBin() string {
	buff := make([]byte, 0)

	for chunks := c.GetI(); chunks > 0; chunks-- {
		buff = append(buff, c.getRaw()...)
	}
	return base64.StdEncoding.EncodeToString(buff)
}

// ------------------------------------------  put

func (c *Connection) PutInt(v int) {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(v))
	c.putData(buf)
}

func (c *Connection) PutLong(v int64) {
	_v := make([]byte, 8)
	binary.BigEndian.PutUint64(_v, uint64(v))
	c.putData(_v)
}

func (c *Connection) PutDbl(v float64) {
	_v := make([]byte, 8)
	binary.BigEndian.PutUint64(_v, math.Float64bits(v))
	c.putData(_v)
}

func (c *Connection) PutStr(v string) {
	c.putData([]byte(v))
}

func (c *Connection) SendBin(v []byte) {
	var chIdx int
	chunks := len(v) / BinChunkSize
	remBytes := len(v) % BinChunkSize
	if remBytes > 0 {
		chunks++
	}

	c.PutInt(chunks)
	c.Send()

	if chunks == 0 {
		return
	}

	for chIdx = 0; chIdx < chunks-1; chIdx++ {
		c.putData(v[chIdx*BinChunkSize : (chIdx+1)*BinChunkSize])
		c.Send()
	}

	c.putData(v[chIdx*BinChunkSize:])
	c.Send()
}

func (c *Connection) PutIP(v net.IP) {
	_v := []byte{4}
	if v4 := v.To4(); v4 == nil {
		_v[0] = 6
		_v = append(_v, v...)
	} else {
		_v = append(_v, v4...)
	}
	c.putData(_v)
}
