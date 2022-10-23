package ch02

import (
	"testing"
)


func TestNewUserName(t *testing.T) {
	userName := NewUserName("kashiwara")
	if userName.value != "kashiwara" {
		t.Errorf("kashiwaraがvalueに入っていない")
	}
}