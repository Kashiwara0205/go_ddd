package repository_mock

import (
	"go_ddd/entity"
	"go_ddd/value_object"
)

type UserRepositoryMock struct{}

func (u *UserRepositoryMock) Exists(name string) bool {
	return name == "name"
}

func (u *UserRepositoryMock) Save(user *entity.User) {}

func (u *UserRepositoryMock) Find(userID *value_object.UserID) *entity.User {
	inputName := "name"
	userName, _ := value_object.NewUserName(&inputName)

	user, _ := entity.NewUser(userName)

	return user
}

func (u *UserRepositoryMock) Delete(user *entity.User) {}
