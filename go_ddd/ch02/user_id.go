package ch02

import (
    "errors"
)

type UserID struct{
	value *string
}

func NewUserID(value *string) (*UserID, error){

	if(value == nil){
		return nil, errors.New("ユーザIDが入力されていません")
	}

	userID := &UserID{ value: value }

	return userID, nil
}