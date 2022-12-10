package urfa

type AttrHeader struct {
	Type uint8  // code of attribute type
	Ver  uint8  // attribute version
	Len  uint16 // attribute total length
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
