package internal

import (
	log "github.com/sirupsen/logrus"
)

func New(lvl log.Level) *log.Logger {
	logger := log.New()
	// logger.SetFormatter(&log.JSONFormatter{})
	logger.SetFormatter(&log.TextFormatter{})
	logger.SetLevel(lvl)
	logger.SetReportCaller(false)
	// safe as we don't log to file.. for now
	logger.SetNoLock()

	return logger
}
