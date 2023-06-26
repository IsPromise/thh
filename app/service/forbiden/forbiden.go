package forbiden

import (
	"thh/app/models/DataReps"
	"time"
)

func Forbidden(key string) bool {
	if DataReps.Get(key) == "" {
		DataReps.Set(key, time.Now().String())
		return false
	}
	return true
}
