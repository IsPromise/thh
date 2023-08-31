package logging

import (
	"fmt"
)

type CronLogging struct {
}

func (itself CronLogging) Printf(format string, args ...interface{}) {
	std().Info(fmt.Sprintf(format, args...))
}

func (itself CronLogging) Info(msg string, keysAndValues ...interface{}) {

}

// Error logs an error condition.
func (itself CronLogging) Error(err error, msg string, keysAndValues ...interface{}) {
	std().Error("error", msg, "err", err, keysAndValues)
}
