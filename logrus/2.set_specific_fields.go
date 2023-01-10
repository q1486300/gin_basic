package main

import "github.com/sirupsen/logrus"

func main() {
	log := logrus.WithField("app", "study").WithField("service", "logrus")

	log = log.WithFields(logrus.Fields{
		"user_id": 21,
		"ip":      "192.168.1.1",
	})

	log.Error("Hello")
}
