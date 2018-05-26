package command

import(

)

type Command struct {
	CommandName string
	CommandDescription string
}

var commands = []Command

func GetCommand(command *Command) Command {
	for i := 0; i < commands.length; i++{
		if command == i {
			return command
		}
	}
}