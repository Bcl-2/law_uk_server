package logs

import (
	log "github.com/sirupsen/logrus"
)

func InitLogger() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.JSONFormatter{})
}