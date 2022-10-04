package main

import (
	"os"

	"github.com/op/go-logging"
)

const (
	// logFormat  = `[ %{time:15:04:05.000} %{color}%{level:.1s} %{shortfile} ] %{color:reset}%{message}`
	// logFormat = `%{time:15:04:05.000} %{shortfile} %{level:.1s}| %{message}`
	logFormat = `%{time:06-01-02 15:04:05.000} %{shortfile} %{level:.1s}| %{message}`
)

var (
	log *logging.Logger
)

func initLogging() {
	log = logging.MustGetLogger("urfa-go")
	formatter := logging.MustStringFormatter(logFormat)
	lb := logging.NewLogBackend(os.Stdout, "", 0)
	lbf := logging.NewBackendFormatter(lb, formatter)
	lbl := logging.AddModuleLevel(lbf)
	logging.SetBackend(lbl)
}
