package config

import (
	log "github.com/sirupsen/logrus"
)

// SetLoggerConfig configures logrus logger
func SetLoggerConfig(cnf Logger) {
	setLevel(cnf.Level)
	setFormatter(cnf.Formatter)
}

// setLevel configures log level
func setLevel(level string) {
	switch level {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	default:
		panic("invalid logger mode")
	}
}

// setFormatter configures log formatter
func setFormatter(formatter string) {
	switch formatter {
	case "text":
		log.SetFormatter(&log.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		})
	default:
		panic("invelid logger formatter")
	}
}
