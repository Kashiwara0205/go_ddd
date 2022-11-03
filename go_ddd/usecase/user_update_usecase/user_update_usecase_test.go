package user_update_usecase

import (
	. "go_ddd/command"
	. "go_ddd/mocks/repository_mock"
	"testing"
)

func TestHandle(t *testing.T) {
	mock := UserRepositoryMock{}
	usecase := NewUserUpdateUsecase(&mock)
	id := "id1"
	command := NewUserUpdateCommand(&id)

	newUserName := "newName"
	command.SetName(&newUserName)

	newMailAddress := "newMailAddress"
	command.SetMailAddress(&newMailAddress)

	result := usecase.Handle(command)

	if result != nil {
		t.Errorf("ユーザーは正常に削除さています")
	}
}
