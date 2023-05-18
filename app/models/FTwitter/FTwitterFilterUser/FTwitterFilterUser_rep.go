package FTwitterFilterUser

import "github.com/leancodebox/goose/querymaker"

func Create(entity *FTwitterFilterUser) int64 {
	result := builder().Create(entity)
	return result.RowsAffected
}

func save(entity *FTwitterFilterUser) int64 {
	result := builder().Save(entity)
	return result.RowsAffected
}

func saveAll(entities []*FTwitterFilterUser) int64 {
	result := builder().Save(entities)
	return result.RowsAffected
}

func DeleteEntity(entity []*FTwitterFilterUser) int64 {
	result := builder().Delete(entity)
	return result.RowsAffected
}

func Get(id any) (entity FTwitterFilterUser) {
	builder().Where(pid, id).First(&entity)
	return
}

func All() (entities []*FTwitterFilterUser) {
	builder().Find(&entities)
	return
}


func GetWithDeleted(screenName string) (entity FTwitterFilterUser) {
	builder().Unscoped().Where(querymaker.Eq(fieldScreenName, screenName)).First(&entity)
	return
}

func GetByScreenName(screenName string) (entity []*FTwitterFilterUser) {
	builder().Where(querymaker.Eq(fieldScreenName, screenName)).Find(&entity)
	return
}

func Restore(entity *FTwitterFilterUser) int64 {
	return builder().Unscoped().Model(entity).Update(fieldDeletedAt, nil).RowsAffected
}
