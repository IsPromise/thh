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

func Update(entity *Role) {
	builder().Save(entity)
}

func UpdateAll(entities *[]Role) {
	builder().Save(entities)
}

func Delete(entity *Role) int64 {
	result := builder().Delete(entity)
	return result.RowsAffected
}

func Get(id any) (entity Role) {
	builder().Where(pid, id).First(&entity)
	return
}

func GetByName(field, value string) (entity Role) {
	builder().Where(field, value).First(&entity)
	return
}

func All() (entities []Role) {
	builder().Find(&entities)
	return
}

func IsExist(field, value string) bool {
	var count int64
	builder().Where(field+" = ?", value).Count(&count)
	return count > 0
}
