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
	builder().First(&entity, id)
	return
}

func All() (entities []*Entity) {
	builder().Order(querymaker.Desc(fieldCreateTime)).Order(querymaker.Desc(fieldDeadline)).Find(&entities)
	return
}

func QueryAll(needAll bool, status []int) (entities []*Entity) {
	query := builder()
	if len(status) != 0 {
		query.Where(querymaker.In(fieldStatus, status))
	}
	if needAll == false {
		query.Where(querymaker.Ne(fieldStatus, 3))
	}
	query.Order(querymaker.Desc(fieldCreateTime)).Order(querymaker.Desc(fieldDeadline)).Find(&entities)
	return
}
