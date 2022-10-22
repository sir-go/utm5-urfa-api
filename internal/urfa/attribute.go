package urfa

type AttrHeader struct {
	Type uint8
	Ver  uint8
	Len  uint16
}

type Attribute struct {
	Header AttrHeader
	Data   []byte
}

func NewAttribute(attrType uint8) *Attribute {
	return &Attribute{
		Header: AttrHeader{
			Type: attrType,
			Ver:  0,
			Len:  AttrHeaderSize,
		},
	}
}
