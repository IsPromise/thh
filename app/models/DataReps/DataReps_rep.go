package DataReps

import (
	"fmt"
	"thh/bundles/connect/dbconnect"
)

func Set(key string, value string) error {
	dataRep := DataReps{
		Key:   key,
		Value: value,
	}

	// Save the record to the database.
	result := builder().Save(&dataRep)

	// Check for errors and return them if necessary.
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no rows affected")
	}

	return nil
}

func Get(key string) string {
	var dataRep DataReps
	builder().Where(&DataReps{
		Key: key,
	}).First(&dataRep)
	return dataRep.Value
}

func Del(key string) {
	dbconnect.Std().Delete(&DataReps{
		Key: key,
	})
}
