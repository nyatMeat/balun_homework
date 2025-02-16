package compute

type Query struct {
	Cmd  Command
	Args []string
}

func NewQuery(cmd Command, args []string) Query {
	return Query{
		Cmd:  cmd,
		Args: args,
	}
}
