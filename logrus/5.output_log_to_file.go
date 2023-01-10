package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

func main() {
	file, _ := os.OpenFile("logrus/info.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))

	logrus.Info("Hello")
	logrus.Error("出錯了")
	logrus.Errorf("出錯了 %s", "xxx")
	logrus.Errorln("出錯了")

}
