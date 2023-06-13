package ActivityLimitConfig

func Get(id any) (entity ActivityLimitConfig) {
	builder().First(&entity, id)
	return
}

func DeleteEntity(entity *ActivityLimitConfig) {
	builder().Delete(&entity)
}

func Save(entity *ActivityLimitConfig) {
	builder().Save(entity)
}

func Create(entity *ActivityLimitConfig) {
	builder().Create(&entity)

}
