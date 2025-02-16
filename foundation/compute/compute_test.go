package compute

import (
	"balun_homework_1/foundation/logger"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeError(t *testing.T) {
	t.Parallel()

	p := NewStdParser()

	ctx := context.Background()
	c := NewStdCompute(p, logger.CreateMock())

	tests := map[string]struct {
		in  string
		err error
	}{
		"empty request": {
			in:  "",
			err: fmt.Errorf("parse error"),
		},
		"incorrect command": {
			in:  "somecmd",
			err: fmt.Errorf("parse error"),
		},
		"incorrect command with lower case get": {
			in:  "get",
			err: fmt.Errorf("parse error"),
		},
		"incorrect command with lower case set": {
			in:  "set",
			err: fmt.Errorf("parse error"),
		},
		"incorrect command with lower case del": {
			in:  "del",
			err: fmt.Errorf("parse error"),
		},
		"GET: without args": {
			in:  "GET",
			err: fmt.Errorf("parse error"),
		},
		"GET: with 2 args": {
			in:  "GET key value",
			err: fmt.Errorf("parse error"),
		},
		"SET: without args": {
			in:  "SET",
			err: fmt.Errorf("parse error"),
		},
		"SET: with 1 args": {
			in:  "SET key",
			err: fmt.Errorf("parse error"),
		},
		"SET: with 3 args": {
			in:  "SET v v v",
			err: fmt.Errorf("parse error"),
		},
		"DEL: without args": {
			in:  "DEL",
			err: fmt.Errorf("parse error"),
		},
		"DEL: with 2 args": {
			in:  "DEL v v v v",
			err: fmt.Errorf("parse error"),
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := c.Handle(ctx, test.in)
			assert.Equal(t, err, test.err)
		})
	}
}

func TestComputeRequestValid(t *testing.T) {
	t.Parallel()

	p := NewStdParser()

	c := NewStdCompute(p, logger.CreateMock())

	ctx := context.Background()

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
			query, _ := c.Handle(ctx, test.in)
			assert.Equal(t, query, test.query)
		})
	}
}
