package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	zlog "github.com/rs/zerolog/log"
)

func main() {
	zlog.Debug().Msg("process command line arguments...")
	fCfgPath := flag.String("c", "config.yml", "path to conf file")
	fAddr := flag.String("h", "localhost:8081", "service address:port")
	flag.Parse()

	zlog.Debug().Str("path", *fCfgPath).Msg("load config file...")
	cfg, err := LoadConfig(*fCfgPath)
	eh(err)

	server := NewServer(cfg.Billings)

	//goland:noinspection HttpUrlsUsage
	zlog.Info().Str("address", fmt.Sprintf("http://%s/api/v1", *fAddr)).Msg("run service")
	http.Handle("/api/v1", server)

	eh((&http.Server{
		Addr:         *fAddr,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}).ListenAndServe())
}
