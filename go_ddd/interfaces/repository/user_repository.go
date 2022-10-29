package repository

import (
	"go_ddd/entity"
)

type UserRepositoryI interface {
	Exists(name string) bool
	Save(user *entity.User)
}
