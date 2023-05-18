package Permission

func Get(id any) (entity Permission) {
	builder().Where(pid, id).First(&entity)
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
