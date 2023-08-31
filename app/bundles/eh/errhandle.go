package eh

import (
	"thh/app/bundles/logging"
)

func PrIF(err error) bool {
	if err != nil {
		logging.Error("err", err)
		return true
	}
	return false
}
