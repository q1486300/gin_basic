package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Error("出錯了")
	logrus.Warn("警告")
	logrus.Info("資訊")
	logrus.Debug("debug")
	logrus.Println("印出")

	fmt.Println(logrus.GetLevel())
}
