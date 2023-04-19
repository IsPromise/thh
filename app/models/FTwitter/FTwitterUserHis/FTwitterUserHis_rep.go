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

func Update(entity *FTwitterUserHis) {
	builder().Save(entity)
}

func UpdateAll(entities *[]FTwitterUserHis) {
	builder().Save(entities)
}

func Delete(entity *FTwitterUserHis) int64 {
	result := builder().Delete(&entity)
	return result.RowsAffected
}

func Get(id any) (entity FTwitterUserHis) {
	builder().Where(pid, id).First(&entity)
	return
}

func GetBy(field, value string) (entity FTwitterUserHis) {
	builder().Where(field+" = ?", value).First(&entity)
	return
}

func All() (entities []FTwitterUserHis) {
	builder().Find(&entities)
	return
}

func IsExist(field, value string) bool {
	var count int64
	builder().Where(field+" = ?", value).Count(&count)
	return count > 0
}
