package main

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init_log() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}
