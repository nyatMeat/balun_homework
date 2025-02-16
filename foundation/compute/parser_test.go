package compute

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseError(t *testing.T) {
	t.Parallel()

	parser := NewStdParser()

	tests := map[string]struct {
		in  string
		err error
	}{
		"empty request": {
			in:  "",
			err: fmt.Errorf("invalid query length"),
		},
		"incorrect command": {
			in:  "somecmd",
			err: fmt.Errorf("unsupported command: somecmd"),
		},
		"incorrect command with lower case get": {
			in:  "get",
			err: fmt.Errorf("unsupported command: get"),
		},
		"incorrect command with lower case set": {
			in:  "set",
			err: fmt.Errorf("unsupported command: set"),
		},
		"incorrect command with lower case del": {
			in:  "del",
			err: fmt.Errorf("unsupported command: del"),
		},
		"GET: without args": {
			in:  "GET",
			err: fmt.Errorf("invalid argument count for GET expected 1 argument, got 0"),
		},
		"GET: with 2 args": {
			in:  "GET key value",
			err: fmt.Errorf("invalid argument count for GET expected 1 argument, got 2"),
		},
		"SET: without args": {
			in:  "SET",
			err: fmt.Errorf("invalid argument count for SET expected 2 arguments, got 0"),
		},
		"SET: with 1 args": {
			in:  "SET key",
			err: fmt.Errorf("invalid argument count for SET expected 2 arguments, got 1"),
		},
		"SET: with 3 args": {
			in:  "SET v v v",
			err: fmt.Errorf("invalid argument count for SET expected 2 arguments, got 3"),
		},
		"DEL: without args": {
			in:  "DEL",
			err: fmt.Errorf("invalid argument count for DEL expected 1 argument, got 0"),
		},
		"DEL: with 2 args": {
			in:  "DEL v v v v",
			err: fmt.Errorf("invalid argument count for DEL expected 1 argument, got 4"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := parser.Parse(test.in)
			assert.Equal(t, err, test.err)
		})
	}
}

func TestParseRequestValid(t *testing.T) {
	t.Parallel()

	parser := NewStdParser()

	tests := map[string]struct {
		in    string
		query Query
	}{
		"correct GET test": {
			in:    "GET key",
			query: Query{Cmd: CommandGet, Args: []string{"key"}},
		},
		"correct SET test": {
			in:    "SET key value",
			query: Query{Cmd: CommandSet, Args: []string{"key", "value"}},
		},
		"correct DEL test": {
			in:    "DEL key",
			query: Query{Cmd: CommandDel, Args: []string{"key"}},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			query, _ := parser.Parse(test.in)
			assert.Equal(t, query, test.query)
		})
	}
}
