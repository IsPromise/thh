package RolePermission

func Get(id any) (entity RolePermission) {
	builder().Where(pid, id).First(&entity)
	return
}

func Delete(entity *RolePermission) {
	builder().Delete(entity)
}

func Save(entity *RolePermission) {
	builder().Save(entity)
}

func Create(entity *RolePermission) {
	builder().Create(&entity)

}
