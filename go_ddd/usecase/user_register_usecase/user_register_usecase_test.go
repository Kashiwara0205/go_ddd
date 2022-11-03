package user_register_usecase

import (
	. "go_ddd/command"
	. "go_ddd/mocks/repository_mock"
	"testing"
)

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
