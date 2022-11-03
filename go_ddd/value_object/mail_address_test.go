package value_object

import (
	"testing"
)

func TestNewMailAddress(t *testing.T) {
	inputUserMailAddress := "MailAddress1"
	userMailAddress, _ := NewMailAddress(&inputUserMailAddress)
	if *userMailAddress.value != "MailAddress1" {
		t.Errorf("MailAddress1がvalueに入っていない")
	}
}

func TestUserMailAddressIsNil(t *testing.T) {
	_, err := NewMailAddress(nil)
	if err.Error() != "メールが入力されていません" {
		t.Errorf("意図していないエラーメッセージ: %s", err.Error())
	}
}
