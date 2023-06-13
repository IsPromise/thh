package Permission

func Get(id any) (entity Permission) {
	builder().First(&entity, id)
	return
}

func Delete(entity *Permission) {
	builder().Delete(&entity)
}

func Save(entity *Permission) {
	builder().Save(entity)
}

func Create(entity *Permission) {
	builder().Create(&entity)

}
