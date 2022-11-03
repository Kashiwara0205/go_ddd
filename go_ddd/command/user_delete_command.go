package command

type UserDeleteCommand struct {
	id *string
}

func NewUserDeleteCommand(id *string) *UserDeleteCommand {
	command := UserDeleteCommand{
		id: id,
	}

	return &command
}

func (u *UserDeleteCommand) ID() *string {
	return u.id
}
