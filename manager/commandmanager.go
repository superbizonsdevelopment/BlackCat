package command

var(
	Commands = make(map[string]Command)
	CommandList = []Command {
		Command {
			Name: "!help",
			Description: "help command description",
			MinArgs: 0,
			MaxArgs: 0,
			Permission: 0,
		},
		Command {
			Name: "!avatar",
			Description: "avatar command description",
			MinArgs: 0,
			MaxArgs: 0,
			Permission: 0,
		},
	}
)

func RegisterCommands() {
	for _, cmd := range CommandList{
		Commands[cmd.Name] = Command {
			cmd.Name,
			cmd.Description,
			cmd.MinArgs,
			cmd.MaxArgs,
			cmd.Permission,
		}
	}
}