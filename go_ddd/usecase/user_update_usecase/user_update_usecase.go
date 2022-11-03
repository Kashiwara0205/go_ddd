package user_update_usecase

import (
	"errors"
	. "go_ddd/command"
	. "go_ddd/interfaces/repository"
	. "go_ddd/service"
	. "go_ddd/value_object"
)

type UserUpdateUsecase struct {
	userRepository UserRepositoryI
	userService    *UserService
}

func NewUserUpdateUsecase(userRepository UserRepositoryI) *UserUpdateUsecase {
	return &UserUpdateUsecase{
		userRepository: userRepository,
		userService:    NewUserService(userRepository),
	}
}

func (u *UserUpdateUsecase) Handle(command *UserUpdateCommand) error {
	userID, userIDErr := NewUserID(command.ID())

	if userIDErr != nil {
		return userIDErr
	}

	user := u.userRepository.Find(userID)

	if nil == user {
		return errors.New("消すユーザが見つかりませんでした")
	}

	if name := command.Name(); name != nil {

		newUserName, userNameErr := NewUserName(name)

		if userNameErr != nil {
			return userNameErr
		}

		user.ChangeUserName(newUserName)

		if u.userService.Exists(user) {
			return errors.New("ユーザ名が重複しています")
		}

	}

	if mailAddress := command.MailAddress(); mailAddress != nil {
		newMailAddress, mailAddressErr := NewMailAddress(mailAddress)

		if mailAddressErr != nil {
			return mailAddressErr
		}

		user.ChangeMailAddress(newMailAddress)
	}

	u.userRepository.Save(user)

	return nil
}
