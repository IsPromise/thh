package Role

func Create(entity *Role) int64 {
	result := builder().Create(&entity)
	return result.RowsAffected
}

func Save(entity *Role) int64 {
	result := builder().Save(entity)
	return result.RowsAffected
}

func SaveAll(entities *[]Role) int64 {
	result := builder().Save(entities)
	return result.RowsAffected
}

func DeleteEntity(entity *Role) int64 {
	result := builder().Delete(entity)
	return result.RowsAffected
}

func Get(id any) (entity Role) {
	builder().First(&entity, id)
	return
}
