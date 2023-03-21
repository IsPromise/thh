package FTwitterMedia

import querybuild "thh/arms/querymaker"

func Create(entity *FTwitterMedia) int64 {
	result := builder().Create(entity)
	return result.RowsAffected
}

func Save(entity *FTwitterMedia) int64 {
	result := builder().Save(entity)
	return result.RowsAffected
}

func SaveAll(entities *[]FTwitterMedia) int64 {
	result := builder().Save(entities)
	return result.RowsAffected
}

func Update(entity *FTwitterMedia) {
	builder().Save(entity)
}

func UpdateAll(entities *[]FTwitterMedia) {
	builder().Save(entities)
}

func Delete(entity FTwitterMedia) int64 {
	result := builder().Delete(&entity)
	return result.RowsAffected
}

func Get(id any) (entity FTwitterMedia) {
	builder().Where(pid, id).First(&entity)
	return
}

func GetBy(field, value string) (entity FTwitterMedia) {
	builder().Where(field+" = ?", value).First(&entity)
	return
}

func GetByUrl(url string) (entity FTwitterMedia) {
	builder().Where(querybuild.Eq(fieldUrl, url)).First(&entity)
	return
}

func All() (entities []FTwitterMedia) {
	builder().Find(&entities)
	return
}

func IsExist(field, value string) bool {
	var count int64
	builder().Where(field+" = ?", value).Count(&count)
	return count > 0
}
