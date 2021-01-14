package utils

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func NewLogger() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
