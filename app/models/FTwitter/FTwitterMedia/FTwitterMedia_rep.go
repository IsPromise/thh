package FTwitterMedia

import "github.com/leancodebox/goose/querymaker"

func create(entity *FTwitterMedia) int64 {
	result := builder().Create(entity)
	return result.RowsAffected
}

func Save(entity *FTwitterMedia) int64 {
	result := builder().Save(entity)
	return result.RowsAffected
}

func saveAll(entities *[]FTwitterMedia) int64 {
	result := builder().Save(entities)
	return result.RowsAffected
}

func deleteEntity(entity *FTwitterMedia) int64 {
	result := builder().Delete(&entity)
	return result.RowsAffected
}

func get(id any) (entity FTwitterMedia) {
	builder().First(&entity, id)
	return
}

func GetByUrl(url string) (entity FTwitterMedia) {
	builder().Where(querymaker.Eq(fieldUrl, url)).First(&entity)
	return
}

func All() (entities []FTwitterMedia) {
	builder().Find(&entities)
	return
}
