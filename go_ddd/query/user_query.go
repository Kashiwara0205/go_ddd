package Query

import (
	. "go_ddd/entity"
)

type UserQuery struct {
	id          *string
	name        *string
	mailAddress *string
}

func NewUserQuery(source *User) *UserQuery {
	query := UserQuery{
		id:          source.ID(),
		name:        source.Name(),
		mailAddress: source.MailAddress(),
	}

	return &query
}

func (u *UserQuery) ID() *string {
	return u.id
}

func (u *UserQuery) Name() *string {
	return u.name
}

func (u *UserQuery) MailAddress() *string {
	return u.mailAddress
}
