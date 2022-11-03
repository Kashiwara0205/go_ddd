package user_delete_usecase

import (
	. "go_ddd/command"
	. "go_ddd/mocks/repository_mock"
	"testing"
)

func TestHandle(t *testing.T) {
	mock := UserRepositoryMock{}
	usecase := NewUserDeleteUsecase(&mock)
	id := "id1"
	command := NewUserDeleteCommand(&id)
	result := usecase.Handle(command)

	if result != nil {
		t.Errorf("ユーザーは正常に削除さています")
	}
}
