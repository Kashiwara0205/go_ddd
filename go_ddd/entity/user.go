package entity

import (
	"errors"
	"go_ddd/value_object"
)

type User struct {
	id   *value_object.UserID
	name *value_object.UserName
}

func NewUser(name *value_object.UserName) (*User, error) {

	if name == nil {
		return nil, errors.New("ユーザ名が入力されていません")
	}

	user := &User{
		name: name,
	}
	return user, nil
}

func (user *User) Name() string {
	return user.name.Value()
}

func (user *User) Equals(other *User) bool {
	if other == nil {
		return false
	}

	return user.id.Value() == other.id.Value()
}

func (user *User) ChangeUserID(id *value_object.UserID) error {
	if id == nil {
		return errors.New("ユーザIDが入力されていません")
	}

	user.id = id

	return nil
}

func (user *User) ChangeUserName(name *value_object.UserName) error {
	if name == nil {
		return errors.New("ユーザ名が入力されていません")
	}

	user.name = name

	return nil
}
