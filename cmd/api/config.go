package main

import (
	"os"

	"github.com/BurntSushi/toml"
)

type CfgUtmServ struct {
	Addr string `toml:"addr"`
	Cert string `toml:"cert"`
}

type CfgUtm map[string]CfgUtmServ

type CfgService struct {
	Host   string `toml:"host"`
	Port   int    `toml:"port"`
	Secret string `toml:"secret"`
}

type Cfg struct {
	Service CfgService `toml:"service"`
	Utm     CfgUtm     `toml:"utm"`
}

func LoadConfig(confpath string) (*Cfg, error) {
	conf := new(Cfg)
	file, err := os.Open(confpath)
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		return nil, err
	}

	_, err = toml.DecodeFile(confpath, &conf)
	if err != nil {
		return nil, err
	}

	return conf, err
}
