package service

import (
	"go_ddd/entity"
	. "go_ddd/interfaces/repository"
)

type UserService struct {
	userRepository UserRepositoryI
}

func NewUserService(userRepository UserRepositoryI) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) Exists(user *entity.User) bool {

	name := user.Name()

	if nil == user.Name() {
		return false
	}

	return u.userRepository.Exists(*name)
}
