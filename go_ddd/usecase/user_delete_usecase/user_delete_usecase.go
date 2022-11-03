package user_delete_usecase

import (
	"errors"
	. "go_ddd/command"
	. "go_ddd/interfaces/repository"
	. "go_ddd/value_object"
)

type UserDeleteUsecase struct {
	userRepository UserRepositoryI
}

func NewUserDeleteUsecase(userRepository UserRepositoryI) *UserDeleteUsecase {
	return &UserDeleteUsecase{
		userRepository: userRepository,
	}
}

func (u *UserDeleteUsecase) Handle(command *UserDeleteCommand) error {
	userID, userIDErr := NewUserID(command.ID())

	if userIDErr != nil {
		return userIDErr
	}

	user := u.userRepository.Find(userID)

	if nil == user {
		return errors.New("消すユーザが見つかりませんでした")
	}

	u.userRepository.Delete(user)

	return nil
}
