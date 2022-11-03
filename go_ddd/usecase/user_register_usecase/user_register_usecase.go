package user_register_usecase

import (
	"errors"
	. "go_ddd/command"
	"go_ddd/entity"
	. "go_ddd/interfaces/repository"
	. "go_ddd/service"
	"go_ddd/value_object"
)

type UserRegisterUsecase struct {
	userRepository UserRepositoryI
	userService    *UserService
}

func NewUserRegisterUsecase(userRepository UserRepositoryI) *UserRegisterUsecase {
	return &UserRegisterUsecase{
		userRepository: userRepository,
		userService:    NewUserService(userRepository),
	}
}

func (u *UserRegisterUsecase) Handle(command *UserRegisterCommand) error {

	userNameValueObject, userNameErr := value_object.NewUserName(command.Name())

	if userNameErr != nil {
		return userNameErr
	}

	user, userEntityErr := entity.NewUser(userNameValueObject)

	if userEntityErr != nil {
		return userEntityErr
	}

	if u.userService.Exists(user) {
		return errors.New("このユーザ名は既に存在します")
	}

	u.userRepository.Save(user)

	return nil
}
