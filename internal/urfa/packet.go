package urfa

type PacketType uint8

type PacketHeader struct {
	Type uint8
	Ver  uint8
	Len  uint16
}

type Packet struct {
	Header     PacketHeader
	Attributes []Attribute
}

func NewPacket(packetType uint8) *Packet {
	return &Packet{
		Header: PacketHeader{
			Type: packetType,
			Ver:  PacketVersion,
			Len:  PacketHeaderSize,
		},
	}
}

func (pkt *Packet) GetAttr(attrType uint8) (attr *Attribute) {
	for _, attr := range pkt.Attributes {
		if attr.Header.Type == attrType {
			return &attr
		}
	}
	return nil
}

func (pkt *Packet) AddAttr(attrType uint8, data []byte) {
	attr := *NewAttribute(attrType)
	attr.Header.Len += uint16(len(data))
	attr.Data = data
	pkt.Header.Len += attr.Header.Len
	pkt.Attributes = append(pkt.Attributes, attr)
}
