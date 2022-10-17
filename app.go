package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/darkmane/traveller/handlers"

	"github.com/darkmane/traveller/util"
)

func main() {
	cfg := util.GetConfig()
	switch strings.ToLower(cfg.LogLevel) {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
		break
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
		break
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		break
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		break
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		break
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		break
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
		break
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	log.Info().Msg(fmt.Sprintf("New Database: %v, Username: %v, Seed: %v", cfg.Database, cfg.Username, cfg.Seed))
	handlers.RegisterHandlers(http.HandleFunc)
	http.ListenAndServe(":80", nil)
}
