package Users

import (
	"github.com/leancodebox/goose/querymaker"
	"thh/app/bundles/algorithm"
)

func Get(id any) (entity Users, err error) {
	err = builder().First(&entity, id).Error
	return
}

func Verify(username string, password string) (*Users, error) {
	var user Users
	err := builder().Where(querymaker.Eq(fieldUsername, username)).First(&user).Error
	if err != nil {
		return &user, err
	}
	err = algorithm.VerifyEncryptPassword(user.Password, password)
	if err != nil {
		return &Users{}, err
	}
	return &user, nil
}

func MakeUser(name string, password string, email string) *Users {
	user := Users{Username: name, Email: email}
	user.SetPassword(password)
	return &user
}

func Create(entity *Users) error {
	return builder().Create(&entity).Error
}

func Save(entity *Users) error {
	return builder().Save(entity).Error
}

func All() (entities []*Users) {
	builder().Find(&entities)
	return
}

func GetByUsername(username string) (entities *Users) {
	builder().Where(querymaker.Eq(fieldUsername, username)).First(entities)
	return
}

func Delete(user *Users) {
	builder().Delete(user)
}
func UnscopedDelete(user *Users) {
	builder().Unscoped().Delete(user)
}

func UnscopedGet(id any) (entity Users, err error) {
	err = builder().Unscoped().First(&entity, id).Error
	return
}
