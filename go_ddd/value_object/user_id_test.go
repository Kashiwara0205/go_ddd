package value_object

import (
	"testing"
)

func TestNewUserID(t *testing.T) {
	inputUserID := "ID1"
	userID, _ := NewUserID(&inputUserID)
	if *userID.value != "ID1" {
		t.Errorf("ID1がvalueに入っていない")
	}
}

func TestUserIDIsNil(t *testing.T) {
	_, err := NewUserID(nil)
	if err.Error() != "ユーザIDが入力されていません" {
		t.Errorf("意図していないエラーメッセージ: %s", err.Error())
	}
}
