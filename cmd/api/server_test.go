package main

import (
	"testing"
)

func Test_getMethod(t *testing.T) {
	tests := []struct {
		name      string
		reqObj    RpcReq
		wantBName string
		wantMName string
		wantOk    bool
	}{
		{"empty", RpcReq{Method: ""}, "", "", false},
		{"ok", RpcReq{Method: "bil.method"}, "bil", "method", true},
		{"ok", RpcReq{Method: "bil.meth.od"}, "bil", "meth", true},
		{"bad", RpcReq{Method: ".meth.od"}, "", "", false},
		{"bad", RpcReq{Method: "meth."}, "", "", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBName, gotMName, gotOk := getMethod(&tt.reqObj)
			if gotBName != tt.wantBName {
				t.Errorf("getMethod() gotBName = %v, want %v", gotBName, tt.wantBName)
			}
			if gotMName != tt.wantMName {
				t.Errorf("getMethod() gotMName = %v, want %v", gotMName, tt.wantMName)
			}
			if gotOk != tt.wantOk {
				t.Errorf("getMethod() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
