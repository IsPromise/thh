package FTwitterUserHis

func Create(entity *FTwitterUserHis) int64 {
	result := builder().Create(entity)
	return result.RowsAffected
}

func Save(entity *FTwitterUserHis) int64 {
	result := builder().Save(entity)
	return result.RowsAffected
}

func SaveAll(entities *[]FTwitterUserHis) int64 {
	result := builder().Save(entities)
	return result.RowsAffected
}
func DeleteEntity(entity *FTwitterUserHis) int64 {
	result := builder().Delete(&entity)
	return result.RowsAffected
}

func Get(id any) (entity FTwitterUserHis) {
	builder().First(&entity, id)
	return
}

func All() (entities []FTwitterUserHis) {
	builder().Find(&entities)
	return
}
