package command

type UserRegisterCommand struct {
	name *string
}

func CreateUserRegisterCommand(name *string) UserRegisterCommand {
	command := UserRegisterCommand{
		name: name,
	}

	return command
}

func (u UserRegisterCommand) Name() *string {
	return u.name
}
