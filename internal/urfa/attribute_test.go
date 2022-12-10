package urfa

import (
	"reflect"
	"testing"
)

func TestNewAttribute(t *testing.T) {
	type args struct {
		attrType uint8
	}
	tests := []struct {
		name string
		args args
		want *Attribute
	}{
		{"any",
			args{AttrTypeCall},
			&Attribute{AttrHeader{Type: AttrTypeCall, Ver: 0, Len: AttrHeaderSize}, nil}},
		{"any",
			args{AttrTypeLogin},
			&Attribute{AttrHeader{Type: AttrTypeLogin, Ver: 0, Len: AttrHeaderSize}, nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewAttribute(tt.args.attrType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAttribute() = %v, want %v", got, tt.want)
			}
		})
	}
}
