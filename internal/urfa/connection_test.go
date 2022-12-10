package urfa

import (
	"testing"
)

func Test_packetTypeIsExpected(t *testing.T) {
	type args struct {
		pType    uint8
		expected []uint8
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty", args{0, nil}, true},
		{"yes", args{4, []uint8{3, 4, 5, 8}}, true},
		{"no", args{44, []uint8{3, 4, 5, 8}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := packetTypeIsExpected(tt.args.pType, tt.args.expected); got != tt.want {
				t.Errorf("packetTypeIsExpected() = %v, want %v", got, tt.want)
			}
		})
	}
}
