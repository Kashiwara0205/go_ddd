package ch02

import (
	"testing"
)

func TestNewUserName(t *testing.T) {
	inputName := "kashiwara"
	userName, _ := NewUserName(&inputName)
	if *userName.value != "kashiwara" {
		t.Errorf("kashiwaraがvalueに入っていない")
	}
}

func TestNameIsNil(t *testing.T) {
	_, err := NewUserName(nil)
	if err.Error() != "ユーザ名が入力されていません" {
		t.Errorf("意図していないエラーメッセージ: %s", err.Error())
	}
}

func TestNameIsLessThanTree(t *testing.T) {
	inputName := "k"
	_, err := NewUserName(&inputName)
	if err.Error() != "ユーザ名は3文字以上です k" {
		t.Errorf("意図していないエラーメッセージ: %s", err.Error())
	}
}
