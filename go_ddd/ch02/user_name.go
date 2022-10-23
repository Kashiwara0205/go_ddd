package ch02

type UserName struct{
	value string
}

func NewUserName(value string) UserName{
	return UserName{
		value: value,
	}
}