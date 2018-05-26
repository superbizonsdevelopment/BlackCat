package command

type Command struct {
	Name string
	Description string
	MinArgs int
	MaxArgs int
	Permission int
}