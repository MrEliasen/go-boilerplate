package logger

import (
	"fmt"
	"os"

	"github.com/placeholder/boiler/pkg/logger/internal"
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

func Logger() *log.Logger {
	if logger != nil {
		return logger
	}

	lvl := log.WarnLevel
	e := os.Getenv("LOG_LEVEL")
	if e != "" {
		l, err := log.ParseLevel(e)
		if err != nil {
			fmt.Printf("invalid log level \"%s\", defaulting to \"warn\"", e)
		} else {
			lvl = l
		}
	}

	logger = internal.New(lvl)
	return logger
}
