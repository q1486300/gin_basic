package main

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	_InitLog("logrus/logs", "Gin")

	logrus.Error("xxx")
}

const (
	_black = iota
	_red
	_green
	_yellow
	_blue
	_purple
	_cyan
	_gray
)

// _LogFormatter 日誌自定義格式
type _LogFormatter struct {
	appName string
}

func (l _LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.ErrorLevel:
		levelColor = _red
	case logrus.WarnLevel:
		levelColor = _yellow
	case logrus.InfoLevel:
		levelColor = _blue
	case logrus.DebugLevel:
		levelColor = _cyan
	default:
		levelColor = _gray
	}

	timeFormat := time.Now().Format("2006-01-02 15:04:05")

	var msg string
	if entry.HasCaller() {
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		funcVal := entry.Caller.Function
		msg = fmt.Sprintf("[%s] {%s} \033[3%dm[%s]\033[0m %s [\033[4%dm%s\033[0m] %s\n", l.appName, timeFormat,
			levelColor, strings.ToUpper(entry.Level.String()), fileVal, _green, funcVal, entry.Message)
	} else {
		msg = fmt.Sprintf("[%s] {%s} \033[3%dm[%s]\033[0m %s\n", l.appName, timeFormat,
			levelColor, strings.ToUpper(entry.Level.String()), entry.Message)
	}

	return []byte(msg), nil
}

type LogFileWriter struct {
	file     *os.File
	logPath  string
	fileDate string // 判斷日期切換資料夾
	appName  string
}

func (l LogFileWriter) Write(data []byte) (n int, err error) {
	if l.file == nil {
		return 0, errors.New("file not opened")
	}
	date := time.Now().Format("2006-01-02")
	if l.fileDate != date {
		l.file.Close()
		err := os.MkdirAll(fmt.Sprintf("%s/%s", l.logPath, date), 0755)
		if err != nil {
			return 0, err
		}

		fileName := fmt.Sprintf("%s/%s/%s.log", l.logPath, date, l.appName)
		l.file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return 0, err
		}
		l.fileDate = date
	}
	return l.file.Write(data)
}

func _InitLog(logPath, appName string) {
	fileDate := time.Now().Format("2006-01-02")
	// 創建資料夾
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), 0755)
	if err != nil {
		logrus.Error(err)
		return
	}

	fileName := fmt.Sprintf("%s/%s/%s.log", logPath, fileDate, appName)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		logrus.Error(err)
		return
	}

	fileWriter := LogFileWriter{file, logPath, fileDate, appName}
	logrus.SetOutput(io.MultiWriter(os.Stdout, fileWriter))
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetReportCaller(true)
	logrus.SetFormatter(_LogFormatter{appName})
}
