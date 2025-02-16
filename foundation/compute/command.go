package compute

type Command string

func NewCommand(str string) Command {
	return Command(str)
}

var (
	CommandGet = NewCommand("GET")
	CommandSet = NewCommand("SET")
	CommandDel = NewCommand("DEL")
)
