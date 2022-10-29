package entity

import (
	"go_ddd/value_object"
	"testing"
)

func TestUserUserIDIsNil(t *testing.T) {
	inputUserID := "ID1"
	userID, _ := value_object.NewUserID(&inputUserID)

	_, err := NewUser(userID, nil)

	if err.Error() != "ユーザ名が入力されていません" {
		t.Errorf("意図していないエラーメッセージ: %s", err.Error())
	}
}

func TestUserNameIsNil(t *testing.T) {
	inputName := "name"
	userName, _ := value_object.NewUserName(&inputName)

	_, err := NewUser(nil, userName)

	if err.Error() != "ユーザIDが入力されていません" {
		t.Errorf("意図していないエラーメッセージ: %s", err.Error())
	}
}

func TestEquals(t *testing.T) {
	inputUserID := "ID1"
	userID, _ := value_object.NewUserID(&inputUserID)

	inputName := "name"
	userName, _ := value_object.NewUserName(&inputName)

	user, _ := NewUser(userID, userName)

	inputOtherUserID := "ID1"
	otherUserID, _ := value_object.NewUserID(&inputOtherUserID)

	inputOtherName := "name2"
	otherUserName, _ := value_object.NewUserName(&inputOtherName)

	otherUser, _ := NewUser(otherUserID, otherUserName)

	if !user.Equals(otherUser) {
		t.Errorf("ユーザAとユーザBは等しい: %v", user.Equals(otherUser))
	}
}

func TestEqualsWhenDiffrentID(t *testing.T) {
	inputUserID := "ID1"
	userID, _ := value_object.NewUserID(&inputUserID)

	inputName := "name"
	userName, _ := value_object.NewUserName(&inputName)

	user, _ := NewUser(userID, userName)

	inputOtherUserID := "ID2"
	otherUserID, _ := value_object.NewUserID(&inputOtherUserID)

	inputOtherName := "name2"
	otherUserName, _ := value_object.NewUserName(&inputOtherName)

	otherUser, _ := NewUser(otherUserID, otherUserName)

	if user.Equals(otherUser) {
		t.Errorf("ユーザAとユーザBは等しくない: %v", user.Equals(otherUser))
	}
}

func TestEqualsWheOtherUserIsNil(t *testing.T) {
	inputUserID := "ID1"
	userID, _ := value_object.NewUserID(&inputUserID)

	inputName := "name"
	userName, _ := value_object.NewUserName(&inputName)

	user, _ := NewUser(userID, userName)

	if user.Equals(nil) {
		t.Errorf("ユーザBは存在しないのでFalseになる")
	}
}
