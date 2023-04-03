package logger

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"thh/arms"
	"thh/bundles/bootstrap"
	"thh/bundles/config"

	"github.com/sirupsen/logrus"
)

const (
	LogTypeStdout = "stdout"
	LogTypeFile   = "file"
)

var log = logrus.StandardLogger()

func Std() *logrus.Logger {
	return log
}

func Info(args ...any) {
	Std().Info(args...)
}

func Printf(format string, args ...interface{}) {
	Std().Printf(format, args...)
}

func Println(args ...interface{}) {
	Std().Println(args...)
}

func Error(args ...interface{}) {
	Std().Error(args...)
}

func ErrIf(err error) bool {
	if err != nil {
		Std().Error(err)
		return true
	}
	return false
}

func init() {
	bootstrap.AddDInit(Init)
}

func Init() {
	var (
		logType = config.Get("LOG_TYPE")
		logPath = config.Get("LOG_PATH", "./storage/log/app.log")
		debug   = config.GetBool("APP_DEBUG")
	)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetReportCaller(true)
	//logrus.SetFormatter(&TextFormatter{})
	logrus.SetFormatter(&LogFormatter{})
	//logrus.SetFormatter(&logrus.TextFormatter{
	//	DisableQuote: true,
	//	//TimestampFormat: "2006-01-02 15:04:05", //时间格式
	//	FullTimestamp: true,
	//	ForceColors:   conf.LogType() == conf.LogTypeStdout,
	//})

	log.Out = os.Stdout
	if debug {
		log.Level = logrus.TraceLevel
	}

	switch logType {
	case LogTypeStdout:
		return
	case LogTypeFile:
		// You could set this to any `io.Writer` such as a file
		if err := arms.FilePutContents(logPath, []byte(""), true); err != nil {
			log.Info(err)
			return
		}

		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Info("Failed to log to file, using default stderr")
			return
		}
		log.Out = file
		return
	default:
		log.Info("Unknown Log Output Type")
		return
	}

}

type LogFormatter struct{}

func (m *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}

	timestamp := entry.Time.Format("2006-01-02 15:04:05.000")
	var msg string
	//entry.Logger.SetReportCaller(true)
	//HasCaller()为true才会有调用信息
	trace := arms.MyTrace()
	// INFO[2023-01-30T18:46:26+08:00]/Users/thh/workspace/thh/arms/logger/logger.go:22 thh/arms/logger.Info() use sqlite
	if entry.HasCaller() {
		fName := filepath.Base(entry.Caller.File)
		msg = fmt.Sprintf("[%-7s] %v [%s] [%s:%d %s] %s\n",
			timestamp, entry.Level.String(), trace.GetNextTrace(), fName, entry.Caller.Line, entry.Caller.Function, entry.Message)
	} else {
		msg = fmt.Sprintf("[%v] [%s] [%s] %s\n", trace.GetNextTrace(), timestamp, entry.Level, entry.Message)
	}

	b.WriteString(msg)
	return b.Bytes(), nil
}
