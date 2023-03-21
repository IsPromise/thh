package eh

import (
	"fmt"
	"thh/arms/logger"
	"thh/bundles/bootstrap"
)

var logManager loggerManager

type loggerManager struct {
	setLogger bool
	logger    Logger
}

func (itself *loggerManager) log(err error) {
	if itself.setLogger {
		itself.logger.Error(err)
	} else {
		fmt.Println(err)
	}
}

type Logger interface {
	Error(...any)
}

func init() {
	bootstrap.AddDInit(bundleInit)
}

func bundleInit() {
	fmt.Println("init eh")
	InitLogger(logger.Std())
}

func InitLogger(logger Logger) {
	logManager.logger = logger
	logManager.setLogger = true
}

func PrIF(err error) bool {
	if err != nil {
		logManager.log(err)
		return true
	}
	return false
}
