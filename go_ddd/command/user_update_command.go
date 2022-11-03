package command

type UserUpdateCommand struct {
	id          *string
	name        *string
	mailAddress *string
}

func CreateUserUpdateCommand(id *string) *UserUpdateCommand {
	command := UserUpdateCommand{
		id: id,
	}

	return &command
}

func (u *UserUpdateCommand) ID() *string {
	return u.id
}

func (u *UserUpdateCommand) Name() *string {
	return u.name
}

func (u *UserUpdateCommand) SetName(name *string) {
	u.name = name
}

func (u *UserUpdateCommand) MailAddress() *string {
	return u.mailAddress
}

func (u *UserUpdateCommand) SetMailAddress(mailAddress *string) {
	u.mailAddress = mailAddress
}
