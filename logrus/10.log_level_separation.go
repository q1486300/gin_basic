package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
	"time"
)

const (
	_BLACK = iota
	_RED
	_GREEN
	_YELLOW
	_BLUE
	_PURPLE
	_CYAN
	_GRAY

	ALL_LOG  = "all"
	ERR_LOG  = "err"
	WARN_LOG = "warn"
	INFO_LOG = "info"
)

func main() {
	mLog.Error("Hello")
	mLog.Error("err")
	mLog.Warn("warn")
	mLog.Info("info")
	mLog.Println("print")
}

var mLog *logrus.Logger

func init() {
	mLog = newLog("logrus/logs", "Gin")
}

func newLog(logPath, appName string) *logrus.Logger {
	allFile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, ALL_LOG), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	errFile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, ERR_LOG), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	warnFile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, WARN_LOG), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	infoFile, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logPath, INFO_LOG), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		logrus.Error(err)
		return nil
	}

	fileHook := FileLevelHook{allFile, errFile, warnFile, infoFile}

	mLog := logrus.New()
	mLog.AddHook(fileHook)
	mLog.SetReportCaller(true)
	mLog.SetFormatter(LogLevelFormatter{appName})
	mLog.SetLevel(logrus.DebugLevel)
	return mLog
}

type LogLevelFormatter struct {
	appName string
}

func (l LogLevelFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var levelColor int
	switch entry.Level {
	case logrus.ErrorLevel:
		levelColor = _RED
	case logrus.WarnLevel:
		levelColor = _YELLOW
	case logrus.InfoLevel:
		levelColor = _BLUE
	case logrus.DebugLevel:
		levelColor = _CYAN
	default:
		levelColor = _GRAY
	}

	timeFormat := time.Now().Format("2006-01-02 15:04:05")

	var msg string
	if entry.HasCaller() {
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		funcVal := entry.Caller.Function
		msg = fmt.Sprintf("[%s] {%s} \033[3%dm[%s]\033[0m %s [\033[4%dm%s\033[0m] %s\n", l.appName, timeFormat,
			levelColor, strings.ToUpper(entry.Level.String()), fileVal, _GREEN, funcVal, entry.Message)
	} else {
		msg = fmt.Sprintf("[%s] {%s} \033[3%dm[%s]\033[0m %s\n", l.appName, timeFormat,
			levelColor, strings.ToUpper(entry.Level.String()), entry.Message)
	}

	return []byte(msg), nil
}

type FileLevelHook struct {
	file     *os.File
	errFile  *os.File
	warnFile *os.File
	infoFile *os.File
}

func (f FileLevelHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (f FileLevelHook) Fire(entry *logrus.Entry) error {
	line, _ := entry.String()
	switch entry.Level {
	case logrus.ErrorLevel:
		f.errFile.Write([]byte(line))
	case logrus.WarnLevel:
		f.warnFile.Write([]byte(line))
	case logrus.InfoLevel:
		f.infoFile.Write([]byte(line))
	}
	f.file.Write([]byte(line))
	return nil
}
