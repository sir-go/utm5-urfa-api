package urfa

import (
	"bytes"
	"net"
	"testing"
)

func TestConnection_PutInt(t *testing.T) {
	type args struct{ v int }
	tests := []struct {
		name     string
		conn     *Connection
		args     args
		wantData []byte
	}{
		{"empty", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{0}, []byte{0, 0, 0, 0}},
		{"small", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{8}, []byte{0, 0, 0, 8}},
		{"big", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{2 << 30}, []byte{128, 0, 0, 0}},
		{"negative", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{-12345678}, []byte{255, 67, 158, 178}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.conn.PutInt(tt.args.v)
			attr := tt.conn.outPkt.GetAttr(AttrTypeData)
			if !bytes.Equal(attr.Data, tt.wantData) {
				t.Errorf("outPacket has data = %v, want %v", attr.Data, tt.wantData)
			}
		})
	}
}

func TestConnection_PutLong(t *testing.T) {
	type args struct{ v int64 }
	tests := []struct {
		name     string
		conn     *Connection
		args     args
		wantData []byte
	}{
		{"empty", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{0}, []byte{0, 0, 0, 0, 0, 0, 0, 0}},
		{"small", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{8}, []byte{0, 0, 0, 0, 0, 0, 0, 8}},
		{"big", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{123456789e10}, []byte{17, 34, 16, 244, 118, 141, 180, 0}},
		{"negative", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{-987654321e9}, []byte{242, 75, 37, 160, 187, 234, 86, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.conn.PutLong(tt.args.v)
			attr := tt.conn.outPkt.GetAttr(AttrTypeData)
			if !bytes.Equal(attr.Data, tt.wantData) {
				t.Errorf("outPacket has data = %v, want %v", attr.Data, tt.wantData)
			}
		})
	}
}

func TestConnection_PutDbl(t *testing.T) {
	type args struct{ v float64 }
	tests := []struct {
		name     string
		conn     *Connection
		args     args
		wantData []byte
	}{
		{"empty", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{0}, []byte{0, 0, 0, 0, 0, 0, 0, 0}},
		{"small", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{8.4231}, []byte{64, 32, 216, 160, 144, 45, 224, 13}},
		{"big", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{123456789865.654321}, []byte{66, 60, 190, 153, 29, 105, 167, 130}},
		{"negative", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{-123456789865.654321}, []byte{194, 60, 190, 153, 29, 105, 167, 130}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.conn.PutDbl(tt.args.v)
			attr := tt.conn.outPkt.GetAttr(AttrTypeData)
			if !bytes.Equal(attr.Data, tt.wantData) {
				t.Errorf("outPacket has data = %v, want %v", attr.Data, tt.wantData)
			}
		})
	}
}

func TestConnection_PutStr(t *testing.T) {
	type args struct{ v string }
	tests := []struct {
		name     string
		conn     *Connection
		args     args
		wantData []byte
	}{
		{"empty", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{""}, []byte("")},
		{"full", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{"appears to work as expected"}, []byte("appears to work as expected")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.conn.PutStr(tt.args.v)
			attr := tt.conn.outPkt.GetAttr(AttrTypeData)
			if !bytes.Equal(attr.Data, tt.wantData) {
				t.Errorf("outPacket has data = %v, want %v", attr.Data, tt.wantData)
			}
		})
	}
}

func TestConnection_PutIP(t *testing.T) {
	type args struct{ v net.IP }
	tests := []struct {
		name     string
		conn     *Connection
		args     args
		wantData []byte
	}{
		{"empty", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{net.IP{}}, []byte{6}},
		{"v4", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{net.IP{10, 10, 0, 15}}, []byte{4, 10, 10, 0, 15}},
		{"v6", &Connection{outPkt: NewPacket(PacketTypeSessionData)},
			args{net.IP{32, 1, 13, 184, 18, 52, 0, 1, 2, 34, 21, 255, 254, 63, 181, 8}},
			[]byte{6, 32, 1, 13, 184, 18, 52, 0, 1, 2, 34, 21, 255, 254, 63, 181, 8}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.conn.PutIP(tt.args.v)
			attr := tt.conn.outPkt.GetAttr(AttrTypeData)
			if !bytes.Equal(attr.Data, tt.wantData) {
				t.Errorf("outPacket has data = %v, want %v", attr.Data, tt.wantData)
			}
		})
	}
}
