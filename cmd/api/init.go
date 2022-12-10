package main

import (
	"os"
	"os/signal"
	"time"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func eh(err error) {
	if err != nil {
		zlog.Err(err).Msg("panic")
		panic(err)
	}
}

func initLog() {
	zlog.Logger = zlog.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: time.RFC3339,
	})
}

func initInterrupt() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			zlog.Info().Msg("-- stop --")
			os.Exit(137)
		}
	}()
}

func init() {
	initLog()
	initInterrupt()
	zlog.Info().Msg("-- start --")
}
