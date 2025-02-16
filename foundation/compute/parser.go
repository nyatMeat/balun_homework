package compute

import (
	"fmt"
	"strings"
)

const (
	CommandNameIndex      = 0
	CommandArgsStartIndex = 1
	SetCommandArgsCount   = 2
	GetCommandArgsCount   = 1
	DelCommandArgsCount   = 1
)

type StdParser struct{}

func NewStdParser() *StdParser {
	return &StdParser{}
}

func (s *StdParser) Parse(q string) (Query, error) {
	queryFields := strings.Fields(q)

	if len(queryFields) == 0 {
		return Query{}, fmt.Errorf("invalid query length")
	}

	command := Command(queryFields[CommandNameIndex])

	args := queryFields[CommandArgsStartIndex:]

	argsLen := len(args)

	switch command {
	case CommandGet:
		if argsLen != GetCommandArgsCount {
			return Query{}, fmt.Errorf("invalid argument count for %s expected 1 argument, got %d",
				CommandGet, argsLen)
		}
	case CommandSet:
		if argsLen != SetCommandArgsCount {
			return Query{}, fmt.Errorf("invalid argument count for %s expected 2 arguments, got %d",
				CommandSet, argsLen)
		}
	case CommandDel:
		if argsLen != DelCommandArgsCount {
			return Query{}, fmt.Errorf("invalid argument count for %s expected 1 argument, got %d",
				CommandDel, argsLen)
		}

	default:
		return Query{}, fmt.Errorf("unsupported command: %s", command)
	}

	return NewQuery(command, args), nil
}
