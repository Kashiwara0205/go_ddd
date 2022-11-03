package usecase

import (
	. "go_ddd/command"
	"go_ddd/entity"
	"testing"
)

type UserRepositoryMock struct{}

func (u *UserRepositoryMock) Exists(name string) bool {
	return name == "name"
}

func (u *UserRepositoryMock) Save(user *entity.User) {}

func TestHandle(t *testing.T) {
	mock := UserRepositoryMock{}
	usecase := NewUserRegisterUsecase(&mock)
	userName := "kashiwara"
	command := CreateUserRegisterCommand(&userName)
	result := usecase.Handle(command)

	if result != nil {
		t.Errorf("ユーザーは正常に作成さています")
	}
}
