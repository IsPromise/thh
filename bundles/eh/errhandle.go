package eh

import (
	"thh/arms/logger"
)

type Logger interface {
	Error(...any)
}

func PrIF(err error) bool {
	if err != nil {
		logger.Error(err)
		return true
	}
	return false
}
