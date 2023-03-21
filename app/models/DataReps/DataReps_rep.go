package DataReps

import "thh/bundles/dbconnect"

//func Create(entity *DataReps) int64 {
//	result := builder().Create(&entity)
//	return result.RowsAffected
//}
//
//func Save(entity DataReps) int64 {
//	result := builder().Save(&entity)
//	return result.RowsAffected
//}
//
//func Update(entity DataReps) {
//	builder().Save(&entity)
//}
//
//func UpdateAll(entities []DataReps) {
//	builder().Save(&entities)
//}
//
//func Delete(entity DataReps) int64 {
//	result := builder().Delete(&entity)
//	return result.RowsAffected
//}
//
//
//func GetBy(field, value string) (entity DataReps) {
//	builder().Where(field + " = ?", value).First(&entity)
//	return
//}
//
//func All() (entities []DataReps) {
//	builder().Find(&entities)
//	return
//}
//
//func IsExist(field, value string) bool {
//	var count int64
//	builder().Where(field + " = ?", value).Count(&count)
//	return count > 0
//}

func Set(key string, value string) (err error) {
	var dataRep DataReps
	dataRep.Key = key
	dataRep.Value = value
	if err = builder().Where(&DataReps{Key: key}).First(&dataRep).Error; err != nil {
		dataRep.Key = key
		dataRep.Value = value
		if err = builder().Create(&dataRep).Error; err != nil {
			return err
		}
	} else {
		dataRep.Key = key
		dataRep.Value = value
		err = builder().Save(&dataRep).Error
	}
	return nil
}

func Get(key string) string {
	var dataRep DataReps
	dbconnect.Std().Where(&DataReps{
		Key: key,
	}).First(&dataRep)
	return dataRep.Value
}

func Del(key string) {
	dbconnect.Std().Delete(&DataReps{
		Key: key,
	})
}
