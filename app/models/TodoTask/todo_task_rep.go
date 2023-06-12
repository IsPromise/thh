package TodoTask

import (
	"github.com/leancodebox/goose/querymaker"
)

func Create(entity *Entity) int64 {
	result := builder().Create(entity)
	return result.RowsAffected
}

func Save(entity *Entity) int64 {
	result := builder().Save(entity)
	return result.RowsAffected
}

func saveAll(entities []*Entity) int64 {
	result := builder().Save(entities)
	return result.RowsAffected
}

func deleteEntity(entity *Entity) int64 {
	result := builder().Delete(entity)
	return result.RowsAffected
}

func Get(id any) (entity Entity) {
	builder().Where(pid, id).First(entity)
	return
}

func All() (entities []*Entity) {
	builder().Order(querymaker.Desc(fieldCreateTime)).Order(querymaker.Desc(fieldDeadline)).Find(&entities)
	return
}
