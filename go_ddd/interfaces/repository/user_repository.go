package repository

import (
	"go_ddd/entity"
	. "go_ddd/value_object"
)

type UserRepositoryI interface {
	Exists(name string) bool
	Save(user *entity.User)
	Find(userID *UserID) *entity.User
	Delete(user *entity.User)
}
