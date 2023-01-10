package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Error("Hello")
	logrus.Info("Hello")
	logrus.Warn("Hello")
	logrus.Debug("Hello")
	logrus.Println("Hello")
}
