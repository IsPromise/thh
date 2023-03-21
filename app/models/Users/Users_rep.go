package Users

import (
	"thh/arms"
	querybuild "thh/arms/querymaker"
)

func Get(id any) (entity Users, err error) {
	err = builder().Where(pid, id).First(&entity).Error
	return
}

func Verify(username string, password string) (Users, error) {
	var user Users
	err := builder().Where(querybuild.Eq(fieldUsername, username)).First(&user).Error
	if err != nil {
		return user, err
	}
	err = arms.VerifyPassword(user.Password, password)
	if err != nil {
		return Users{}, err
	}
	return user, nil
}

func MakeUser(name string, password string, email string) *Users {
	user := Users{Username: name, Email: email}
	user.SetPassword(password)
	return &user
}

func Create(entity *Users) error {
	return builder().Create(&entity).Error
}

func Save(entity *Users) int64 {
	result := builder().Save(entity)
	return result.RowsAffected
}

func Update(entity *Users) {
	builder().Save(entity)
}

func UpdateAll(entities *[]Users) {
	builder().Save(entities)
}

func Delete(entity *Users) int64 {
	result := builder().Delete(entity)
	return result.RowsAffected
}

func GetBy(field, value string) (entity Users) {
	builder().Where(field+" = ?", value).First(&entity)
	return
}

func All() (entities []Users) {
	builder().Find(&entities)
	return
}

func IsExist(field, value string) bool {
	var count int64
	builder().Where(field+" = ?", value).Count(&count)
	return count > 0
}
