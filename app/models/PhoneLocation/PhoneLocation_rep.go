package PhoneLocation

func Create(entity *PhoneLocation) int64 {
	result := builder().Create(entity)
	return result.RowsAffected
}

func Save(entity *PhoneLocation) int64 {
	result := builder().Save(entity)
	return result.RowsAffected
}

func SaveAll(entities *[]PhoneLocation) int64 {
	result := builder().Save(entities)
	return result.RowsAffected
}

func Update(entity *PhoneLocation) {
	builder().Save(entity)
}

func UpdateAll(entities *[]PhoneLocation) {
	builder().Save(entities)
}

func Delete(entity *PhoneLocation) int64 {
	result := builder().Delete(entity)
	return result.RowsAffected
}

func Get(id any) (entity *PhoneLocation) {
	builder().Where(pid, id).First(entity)
	return
}

func GetBy(field, value string) (entity PhoneLocation) {
	builder().Where(field+" = ?", value).First(&entity)
	return
}

func All() (entities []PhoneLocation) {
	builder().Find(&entities)
	return
}

func IsExist(field, value string) bool {
	var count int64
	builder().Where(field+" = ?", value).Count(&count)
	return count > 0
}
