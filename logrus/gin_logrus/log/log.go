package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"strings"
	"time"
)

const (
	BLACK = iota
	RED
	GREEN
	YELLOW
	BLUE
	PURPLE
	CYAN
	GRAY
)

func InitLog(logPath, appName string) {
	fileDate := time.Now().Format("2006-01-02")
	// 創建資料夾
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileDate), 0755)
	if err != nil {
		logrus.Error(err)
		return
	}

	fileName := fmt.Sprintf("%s/%s/%s.Log", logPath, fileDate, appName)
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		logrus.Error(err)
		return
	}

	fileHook := FileDateHook{file, logPath, fileDate, appName}

	logrus.AddHook(fileHook)
	logrus.SetReportCaller(true) // 取得文件和行號，需要設置為 true
	logrus.SetFormatter(LogFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}

// _LogFormatter 日誌自定義格式
type LogFormatter struct{}

func (l LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if entry.Data["type"] == "builtin" {
		msg := fmt.Sprintf("%s\n", entry.Message)
		return []byte(msg), nil
	}

	var levelColor int
	switch entry.Level {
	case logrus.ErrorLevel:
		levelColor = RED
	case logrus.WarnLevel:
		levelColor = YELLOW
	case logrus.InfoLevel:
		levelColor = BLUE
	case logrus.DebugLevel:
		levelColor = CYAN
	default:
		levelColor = GRAY
	}

	timeFormat := time.Now().Format("2006-01-02 15:04:05")

	var msg string
	if entry.HasCaller() {
		fileVal := fmt.Sprintf("%s:%d", path.Base(entry.Caller.File), entry.Caller.Line)
		funcVal := entry.Caller.Function
		msg = fmt.Sprintf("[GIN] {%s} \033[3%dm[%s]\033[0m %s [\033[4%dm%s\033[0m] %s\n", timeFormat,
			levelColor, strings.ToUpper(entry.Level.String()), fileVal, GREEN, funcVal, entry.Message)
	} else {
		msg = fmt.Sprintf("[GIN] {%s} \033[3%dm[%s]\033[0m %s\n", timeFormat,
			levelColor, strings.ToUpper(entry.Level.String()), entry.Message)
	}

	return []byte(msg), nil
}

type FileDateHook struct {
	file     *os.File
	logPath  string
	fileDate string // 判斷日期切換資料夾
	appName  string
}

func (f FileDateHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (f FileDateHook) Fire(entry *logrus.Entry) error {
	date := entry.Time.Format("2006-01-02")
	if f.fileDate != date {
		f.file.Close()
		err := os.MkdirAll(fmt.Sprintf("%s/%s", f.logPath, date), 0755)
		if err != nil {
			return err
		}

		fileName := fmt.Sprintf("%s/%s/%s.Log", f.logPath, date, f.appName)
		f.file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return err
		}
		f.fileDate = date
	}
	line, _ := entry.String()
	f.file.Write([]byte(line))
	return nil
}
