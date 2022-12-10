package fn

import (
	"net"
	"testing"

	"urfa-go/internal/urfa"
)

func Test_mustHave(t *testing.T) {
	type args struct {
		p    Dict
		name string
	}
	tests := []struct {
		name      string
		args      args
		wantPanic bool
	}{
		{"empty", args{Dict{}, ""}, true},
		{"ok", args{Dict{"k": 4}, "k"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); (r != nil) != tt.wantPanic {
					t.Errorf("got panic '%v', want %v", r, tt.wantPanic)
				}
			}()
			mustHave(tt.args.p, tt.args.name)
		})
	}
}

func Test_forEach(t *testing.T) {
	type args struct {
		c    conn
		p    Dict
		name string
		cb   func(item interface{})
	}
	tests := []struct {
		name       string
		args       args
		wantLength int
	}{
		{"empty", args{&urfa.Connection{}, Dict{}, "", nil}, 0},
		{
			"emptyArray",
			args{
				&urfa.Connection{},
				Dict{"key": []interface{}{}},
				"key",
				func(item interface{}) {},
			}, 0,
		},
		{
			"ok",
			args{
				&urfa.Connection{},
				Dict{"key": []interface{}{3, 8, 0}},
				"key",
				func(item interface{}) {},
			}, 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLength := forEach(tt.args.c, tt.args.p, tt.args.name, tt.args.cb); gotLength != tt.wantLength {
				t.Errorf("forEach() = %v, want %v", gotLength, tt.wantLength)
			}

		})
	}
}

func Test_putI(t *testing.T) {
	type args struct {
		p    Dict
		name string
		def  []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"okVal", args{Dict{"k": 14}, "k", nil}, 14},
		{"okVal", args{Dict{"k": 14.0}, "k", nil}, 14},
		{"okDef", args{Dict{}, "k", []int{20}}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := putI(&urfa.Connection{}, tt.args.p, tt.args.name, tt.args.def...); got != tt.want {
				t.Errorf("putI() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_putL(t *testing.T) {
	type args struct {
		p    Dict
		name string
		def  []int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{"okValInt", args{Dict{"k": 14}, "k", nil}, 14},
		{"okValFloat", args{Dict{"k": 14.0}, "k", nil}, 14},
		{"okDef", args{Dict{}, "k", []int64{20}}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := putL(&urfa.Connection{}, tt.args.p, tt.args.name, tt.args.def...); got != tt.want {
				t.Errorf("putL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_putD(t *testing.T) {
	type args struct {
		p    Dict
		name string
		def  []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{"okValInt", args{Dict{"k": 14}, "k", nil}, 14},
		{"okValFloat", args{Dict{"k": 654.321}, "k", nil}, 654.321},
		{"okDef", args{Dict{}, "k", []float64{123.456}}, 123.456},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := putD(&urfa.Connection{}, tt.args.p, tt.args.name, tt.args.def...); got != tt.want {
				t.Errorf("putD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_putS(t *testing.T) {
	type args struct {
		p    Dict
		name string
		def  []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"okVal", args{Dict{"k": "some"}, "k", nil}, "some"},
		{"okDef", args{Dict{}, "k", []string{"some"}}, "some"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := putS(&urfa.Connection{}, tt.args.p, tt.args.name, tt.args.def...); got != tt.want {
				t.Errorf("putS() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_putA(t *testing.T) {
	type args struct {
		p    Dict
		name string
		def  []string
	}
	tests := []struct {
		name string
		args args
		want net.IP
	}{
		{"okVal", args{Dict{"k": "10.10.2.14"}, "k", nil}, net.IP{10, 10, 2, 14}},
		{"okDef", args{Dict{}, "k", []string{"192.168.10.17"}}, net.IP{192, 168, 10, 17}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := putA(&urfa.Connection{}, tt.args.p, tt.args.name, tt.args.def...); !got.Equal(tt.want) {
				t.Errorf("putA() = %v, want %v", got, tt.want)
			}
		})
	}
}
