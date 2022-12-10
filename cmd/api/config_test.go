package main

import (
	"reflect"
	"testing"
)

func Test_readConfig(t *testing.T) {
	tests := []struct {
		name     string
		confData string
		want     *Config
		wantErr  bool
	}{
		{
			name:     "empty",
			confData: ``,
			want:     &Config{},
			wantErr:  false,
		},
		{
			name: "valid",
			confData: `
billings:
  test:
    addr: 192.168.122.168:11758
    cert: urfa.crt
  tv:
    addr: 192.168.28.135:11758
    cert: urfa1.crt
`,
			want: &Config{Billings: billingsMap{
				"test": {"192.168.122.168:11758", "urfa.crt"},
				"tv":   {"192.168.28.135:11758", "urfa1.crt"},
			}},
			wantErr: false,
		},
		{
			name: "invalid",
			confData: `
billings:
  test:
    addr: 192.168.122.168:11758
    cert: urfa.crt
  tv:
    addr: 192.168.28.135:11758
    cert: urfa1.crt
	wrong: true
`,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readConfig([]byte(tt.confData))
			if (err != nil) != tt.wantErr {
				t.Errorf("readConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
