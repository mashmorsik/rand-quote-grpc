package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var Nlog *zerolog.Logger

func Logger() {
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339},
	).Level(zerolog.TraceLevel).With().Timestamp().Caller().Logger()

	Nlog = &logger
}
