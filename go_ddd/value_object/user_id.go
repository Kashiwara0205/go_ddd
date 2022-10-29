package value_object

import (
	"errors"
)

type UserID struct {
	value *string
}

func NewUserID(value *string) (*UserID, error) {

	if value == nil {
		return nil, errors.New("ユーザIDが入力されていません")
	}

	userID := &UserID{value: value}

	return userID, nil
}

func (u *UserID) ID() string {
	return *u.value
}
