package ActivityConfig

func Get(id any) (entity ActivityConfig) {
	builder().First(&entity,id)
	return
}

func DeleteEntity(entity *ActivityConfig) {
	builder().Delete(entity)
}

func Save(entity *ActivityConfig) {
	builder().Save(entity)
}
func SaveAll(entities *[]ActivityConfig) {
	builder().Save(entities)
}

func Create(entity *ActivityConfig) {
	builder().Create(entity)
}
