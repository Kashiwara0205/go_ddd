package usecase

import (
	"go_ddd/entity"
	"testing"
)

type UserRepositoryMock struct{}

func (u *UserRepositoryMock) Exists(name string) bool {
	return name == "name"
}

func (u *UserRepositoryMock) Save(user *entity.User) {}

func TestCreateUser(t *testing.T) {
	mock := UserRepositoryMock{}
	usecase := NewUserUsecase(&mock)
	userName := "kashiwara"
	result := usecase.CreateUser(&userName)

	if result != nil {
		t.Errorf("ユーザーは正常に作成さています")
	}
}
