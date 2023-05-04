package eh

import (
	"thh/bundles/logging"
)

type Logger interface {
	Error(...any)
}

func PrIF(err error) bool {
	if err != nil {
		logging.Error(err)
		return true
	}
	return false
}
