package urfa

import (
	"reflect"
	"testing"
)

func TestPacket_AddAttr(t *testing.T) {
	tests := []struct {
		name     string
		attrType uint8
		attrData []byte
	}{
		{"empty", 0, nil},
		{"login", AttrTypeLogin, nil},
		{"call", AttrTypeCall, []byte("some attribute data")},
		{"term", AttrTypeTermination, []byte("some attribute data")},
		{"ssl", AttrTypeSSLRequest, []byte("some attribute data")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			attrWant := NewAttribute(tt.attrType)
			attrWant.Header.Len += uint16(len(tt.attrData))
			attrWant.Data = tt.attrData

			pkt := NewPacket(0)
			pkt.AddAttr(tt.attrType, tt.attrData)
			attr := pkt.GetAttr(tt.attrType)

			if !reflect.DeepEqual(attrWant, attr) {
				t.Errorf("packet added attr = %v, want %v", attr, attrWant)
			}
		})
	}
}
