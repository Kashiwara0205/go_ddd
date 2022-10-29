package entity

import (
	"errors"
	"go_ddd/value_object"
)

type User struct {
	id   *value_object.UserID
	name *value_object.UserName
}

func NewUser(id *value_object.UserID, name *value_object.UserName) (*User, error) {
	if id == nil {
		return nil, errors.New("ユーザIDが入力されていません")
	}

	if name == nil {
		return nil, errors.New("ユーザ名が入力されていません")
	}

	user := &User{
		id:   id,
		name: name,
	}
	return user, nil
}

func (user *User) Equals(other *User) bool {
	if other == nil {
		return false
	}

	return user.id.ID() == other.id.ID()
}
