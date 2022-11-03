package service

import (
	"go_ddd/entity"
	. "go_ddd/mocks/repository_mock"
	"go_ddd/value_object"
	"testing"
)

func TestExists(t *testing.T) {
	mock := UserRepositoryMock{}
	service := NewUserService(&mock)
	inputName := "name"
	userName, _ := value_object.NewUserName(&inputName)
	user, _ := entity.NewUser(userName)

	if !service.Exists(user) {
		t.Errorf("userは存在します")
	}
}

func TestUnExists(t *testing.T) {
	mock := UserRepositoryMock{}
	service := NewUserService(&mock)
	inputName := "kashiwara"
	userName, _ := value_object.NewUserName(&inputName)
	user, _ := entity.NewUser(userName)

	if service.Exists(user) {
		t.Errorf("userは存在しません")
	}
}
