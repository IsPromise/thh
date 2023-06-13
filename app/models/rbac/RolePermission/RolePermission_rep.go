package RolePermission

func Get(id any) (entity RolePermission) {
	builder().First(&entity, id)
	return
}

func DeleteEntity(entity *RolePermission) {
	builder().Delete(entity)
}

func Save(entity *RolePermission) {
	builder().Save(entity)
}

func Create(entity *RolePermission) {
	builder().Create(&entity)
}
