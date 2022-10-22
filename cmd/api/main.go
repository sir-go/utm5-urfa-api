package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

func check(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func initInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Info("-- stop --")
			os.Exit(137)
		}
	}()
}

func init() {
	initLogging()
	initInterrupt()
}

func main() {
	log.Info("-- start --")
	log.Debug("process command line arguments...")
	fCfgPath := flag.String("c", "conf.toml", "path to conf file")
	flag.Parse()

	log.Debug("parse config file")
	cfg, err := LoadConfig(*fCfgPath)
	check(err)

	addr := fmt.Sprintf("%s:%d", cfg.Service.Host, cfg.Service.Port)

	server := NewApiServer(cfg.Utm)

	log.Debugf("http://%s/api/v1", addr)

	http.Handle("/api/v1", server)
	log.Fatal(http.ListenAndServe(addr, nil))
}
