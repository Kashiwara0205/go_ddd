package user_get_usecase

import (
	. "go_ddd/command"
	. "go_ddd/mocks/repository_mock"
	"testing"
)

func TestHandle(t *testing.T) {
	mock := UserRepositoryMock{}
	usecase := NewUserGetUsecase(&mock)
	id := "1"
	command := NewUserGetCommand(&id)
	query, err := usecase.Handle(command)

	if err != nil {
		t.Errorf("ユーザーは正常に削除さています")
	}

	if *query.ID() != "1" {
		t.Errorf("ID=1ではありません")
	}

	if *query.Name() != "name" {
		t.Errorf("Name=nameではありません")
	}

	if query.MailAddress() != nil {
		t.Errorf("MailAddress=nilではありません")
	}
}
