package util

import (
	"os"

	"github.com/rs/zerolog"
)

var logger zerolog.Logger

func Logger() *zerolog.Logger {
	if logger == nil {
		logger = zerolog.New(os.Stdout).With().Timestamp().Logger()
	}
	return &logger
}
