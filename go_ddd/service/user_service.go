package service

import (
	"go_ddd/entity"
)

type userRepositoryI interface {
	Find(name string) bool
	Save(user *entity.User)
}

type UserService struct {
	userRepository userRepositoryI
}

func NewUserService(userRepository userRepositoryI) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (u *UserService) Exists(user *entity.User) bool {
	return u.userRepository.Find(user.Name())
}
