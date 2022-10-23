package ch02

import (
    "errors"
	"fmt"
)

type UserName struct{
	value *string
}

func NewUserName(value *string) (*UserName, error){

	if(value == nil){
		return nil, errors.New("ユーザ名が入力されていません")
	}

	if len(*value)< 3{
		errorMessage := fmt.Sprintf("ユーザ名は3文字以上です %s", *value)
		return nil, errors.New(errorMessage)
	}

	userName := &UserName{ value: value }

	return userName, nil
}