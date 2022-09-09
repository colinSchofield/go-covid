package config

import (
	"os"

	log "github.com/sirupsen/logrus"
)

// The init function initialises the logger only once
var logger *log.Logger

func init() {
	logger = &log.Logger{
		Out:       os.Stdout,
		Level:     log.InfoLevel,
		Formatter: &log.JSONFormatter{},
	}
}

func Logger() *log.Logger {
	return logger
}
