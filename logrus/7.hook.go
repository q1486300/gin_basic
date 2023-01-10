package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

type MyHook struct {
	Writer io.Writer
}

func (m MyHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel}
}

func (m MyHook) Fire(entry *logrus.Entry) error {
	//entry.Data["app"] = "Gin_basic"
	line, err := entry.String()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read entry, %v", err)
		return err
	}
	m.Writer.Write([]byte(line))
	return nil
}

func main() {
	file, _ := os.OpenFile("logrus/err.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	hook := MyHook{Writer: file}
	logrus.AddHook(hook)

	logrus.Warn("Hello")
	logrus.Error("Hello")
}
