package usecase

import (
	"errors"
	"go_ddd/entity"
	. "go_ddd/interfaces/repository"
	"go_ddd/service"
	"go_ddd/value_object"
)

type UserUsecase struct {
	userRepository UserRepositoryI
}

func NewUserUsecase(userRepository UserRepositoryI) *UserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (u *UserUsecase) CreateUser(userName *string) error {

	userNameValueObject, userNameErr := value_object.NewUserName(userName)

	if userNameErr != nil {
		return userNameErr
	}

	user, userEntityErr := entity.NewUser(userNameValueObject)

	if userEntityErr != nil {
		return userEntityErr
	}

	userService := service.NewUserService(u.userRepository)

	if userService.Exists(user) {
		return errors.New("このユーザ名は既に存在します")
	}

	u.userRepository.Save(user)

	return nil
}
