package entity

import (
	"go_ddd/value_object"
	"testing"
)

func TestUserUserIDIsNil(t *testing.T) {
	_, err := NewUser(nil)

	if err.Error() != "ユーザ名が入力されていません" {
		t.Errorf("意図していないエラーメッセージ: %s", err.Error())
	}
}

func TestEquals(t *testing.T) {
	inputName := "name"
	userName, _ := value_object.NewUserName(&inputName)

	user, _ := NewUser(userName)

	inputID := "ID1"
	userID, _ := value_object.NewUserID(&inputID)
	user.ChangeUserID(userID)

	inputOtherName := "name"
	otherUserName, _ := value_object.NewUserName(&inputOtherName)

	otherUser, _ := NewUser(otherUserName)

	inputOtherID := "ID1"
	otherUserID, _ := value_object.NewUserID(&inputOtherID)
	otherUser.ChangeUserID(otherUserID)

	if !user.Equals(otherUser) {
		t.Errorf("ユーザAとユーザBは等しい: %v", user.Equals(otherUser))
	}
}

func TestEqualsWhenDiffrentID(t *testing.T) {

	inputName := "name"
	userName, _ := value_object.NewUserName(&inputName)

	user, _ := NewUser(userName)

	inputID := "ID1"
	userID, _ := value_object.NewUserID(&inputID)
	user.ChangeUserID(userID)

	inputOtherName := "name2"
	otherUserName, _ := value_object.NewUserName(&inputOtherName)

	otherUser, _ := NewUser(otherUserName)

	inputOtherID := "ID2"
	otherUserID, _ := value_object.NewUserID(&inputOtherID)
	otherUser.ChangeUserID(otherUserID)

	if user.Equals(otherUser) {
		t.Errorf("ユーザAとユーザBは等しくない: %v", user.Equals(otherUser))
	}
}

func TestEqualsWheOtherUserIsNil(t *testing.T) {

	inputName := "name"
	userName, _ := value_object.NewUserName(&inputName)

	user, _ := NewUser(userName)

	if user.Equals(nil) {
		t.Errorf("ユーザBは存在しないのでFalseになる")
	}
}

func TestChangeUserName(t *testing.T) {

	inputName := "name"
	userName, _ := value_object.NewUserName(&inputName)

	user, _ := NewUser(userName)

	changeName := "changeName"
	changeUserName, _ := value_object.NewUserName(&changeName)

	user.ChangeUserName(changeUserName)

	if user.name.Value() != "changeName" {
		t.Errorf("名前が変更されていない")
	}
}

func TestChangeUserNameWhenNameIsNil(t *testing.T) {
	inputName := "name"
	userName, _ := value_object.NewUserName(&inputName)

	user, _ := NewUser(userName)

	err := user.ChangeUserName(nil)

	if err.Error() != "ユーザ名が入力されていません" {
		t.Errorf("意図していないエラーメッセージ: %s", err.Error())
	}
}
