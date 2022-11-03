package user_get_usecase

import (
	. "go_ddd/command"
	. "go_ddd/interfaces/repository"
	. "go_ddd/query"
	. "go_ddd/value_object"
)

type UserGetUsecase struct {
	userRepository UserRepositoryI
}

func NewUserGetUsecase(userRepository UserRepositoryI) *UserGetUsecase {
	return &UserGetUsecase{
		userRepository: userRepository,
	}
}

func (u *UserGetUsecase) Handle(command *UserGetCommand) (*UserQuery, error) {

	targetID, userIDErr := NewUserID(command.ID())

	if userIDErr != nil {
		return nil, userIDErr
	}

	user := u.userRepository.Find(targetID)

	if user == nil {
		return nil, nil
	}

	return NewUserQuery(user), nil
}
