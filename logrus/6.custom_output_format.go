package main

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
)

const (
	colorBlack = iota
	colorRed
	colorGreen
	colorYellow
	colorBlue
	colorPurple
	colorCyan
	colorGray
)

type MyFormatter struct {
	Prefix     string
	TimeFormat string
}

func (m MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 設置顏色
	var color int
	switch entry.Level {
	case logrus.ErrorLevel:
		color = colorRed
	case logrus.WarnLevel:
		color = colorYellow
	case logrus.InfoLevel:
		color = colorBlue
	case logrus.DebugLevel:
		color = colorCyan
	default:
		color = colorGray
	}

	// 設置 buffer
	var b *bytes.Buffer
	if entry.Buffer == nil {
		b = &bytes.Buffer{}
	} else {
		b = entry.Buffer
	}

	// 時間格式化
	formatTime := entry.Time.Format(m.TimeFormat)

	if entry.HasCaller() {
		// 函數名稱
		funcVal := entry.Caller.Function
		// 文件和行號
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)

		// 設置格式
		fmt.Fprintf(b, "[%s] {%s} \033[3%dm[%s]\033[0m %s [\033[4%dm%s\033[0m] %s\n", m.Prefix, formatTime,
			color, strings.ToUpper(entry.Level.String()), fileVal, colorGreen, funcVal, entry.Message)
	} else {
		// 設置格式
		fmt.Fprintf(b, "[%s] {%s} \033[3%dm[%s]\033[0m %s\n", m.Prefix, formatTime,
			color, strings.ToUpper(entry.Level.String()), entry.Message)
	}

	return b.Bytes(), nil
}

var log *logrus.Logger

func init() {
	log = NewLog()
}

func NewLog() *logrus.Logger {
	mLog := logrus.New()
	mLog.SetOutput(os.Stdout)
	mLog.SetReportCaller(true) // 取得文件和行號，需要設置為 true
	mLog.SetFormatter(MyFormatter{Prefix: "GIN", TimeFormat: "2006-01-02 15:04:05"})
	mLog.SetLevel(logrus.DebugLevel)
	return mLog
}

func main() {
	log.Info("Hello")
	log.Error("Hello")
	log.Warn("Hello")
	log.Debug("Hello")
}
