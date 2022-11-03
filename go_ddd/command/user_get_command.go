package command

type UserGetCommand struct {
	id *string
}

func NewUserGetCommand(id *string) *UserGetCommand {
	command := UserGetCommand{
		id: id,
	}

	return &command
}

func (u *UserGetCommand) ID() *string {
	return u.id
}
