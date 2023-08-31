package logging

import (
	"context"
	"fmt"
	"github.com/leancodebox/goose/fileopt"
	"github.com/leancodebox/goose/preferences"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

const (
	LogTypeStdout = "stdout"
	LogTypeFile   = "file"
)

type Entry struct {
	level   slog.Level
	message string
	args    []any
}

var (
	log        *slog.Logger
	logChannel = make(chan *Entry, 1024*512)
	wg         sync.WaitGroup
)

func std() *slog.Logger {
	return log
}

func Info(args ...interface{}) {
	sendLog(slog.LevelInfo, fmt.Sprint(args...))
}

func Printf(format string, args ...interface{}) {
	sendLog(slog.LevelInfo, fmt.Sprintf(format, args...))
}

func Println(args ...interface{}) {
	sendLog(slog.LevelInfo, fmt.Sprintln(args...))
}

func Error(args ...interface{}) {
	sendLog(slog.LevelError, fmt.Sprint(args...))
}

func ErrIf(err error) bool {
	if err != nil {
		sendLog(slog.LevelError, err.Error())
		return true
	}
	return false
}

// Send log entry to the channel
func sendLog(level slog.Level, msg string, args ...any) {

	if caller, success := getCaller(3); success {
		msg = caller + ":" + msg
	}

	entry := &Entry{
		level:   level,
		message: msg,
	}
	logChannel <- entry
}

func getCaller(depth int) (string, bool) {
	pc, file, line, ok := runtime.Caller(depth) // 1 表示跳过当前函数的调用帧
	if ok {
		f := runtime.FuncForPC(pc)
		if f != nil {
			funcName := f.Name()
			return fmt.Sprintf("[%s:%s:%d] message", funcName, filepath.Base(file), line), true
		}
	}
	return "", false
}

func processLogEntries() {
	defer wg.Done()
	for entry := range logChannel {
		std().Log(context.Background(), entry.level, entry.message, entry.args...)
	}
}

var (
	logType = preferences.Get("log.type")
	logPath = preferences.Get("log.path", "./storage/logs/run.log")
	debug   = preferences.GetBool("app.debug", true)
)

func init() {
	var err error
	var logOut io.Writer
	logOut = os.Stdout
	switch logType {
	default:
		slog.Info("Unknown Log Output Type")
	case LogTypeStdout:
	case LogTypeFile:
		if err = fileopt.FilePutContents(logPath, []byte(""), true); err != nil {
			slog.Info("Create log file error", "err", err)
			return
		}
		logOut, err = os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			slog.Info("Failed to log to file, using default stderr", "err", err)
			return
		}
	}

	log = slog.New(slog.NewJSONHandler(logOut, &slog.HandlerOptions{
		AddSource:   true,
		ReplaceAttr: replace,
	}))

	wg.Add(1)
	go processLogEntries()
}

func replace(groups []string, a slog.Attr) slog.Attr {
	// Remove time.
	if a.Key == slog.TimeKey && len(groups) == 0 {
		return slog.Attr{}
	}
	// Remove the directory from the source's filename.
	if a.Key == slog.SourceKey {
		source := a.Value.Any().(*slog.Source)
		source.File = filepath.Base(source.File)
	}
	return a
}

func Shutdown() {
	close(logChannel)
	wg.Wait()
	fmt.Println("logging 👋")
}
