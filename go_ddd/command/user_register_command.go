package command

type UserRegisterCommand struct {
	name *string
}

func NewUserRegisterCommand(name *string) *UserRegisterCommand {
	command := UserRegisterCommand{
		name: name,
	}

	return &command
}

func (u *UserRegisterCommand) Name() *string {
	return u.name
}
