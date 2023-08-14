package eh

import (
	"thh/app/bundles/logging"
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
