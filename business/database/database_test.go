package database

import (
	"balun_homework_1/foundation/compute"
	"balun_homework_1/foundation/logger"
	"balun_homework_1/foundation/storage"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

type DatabaseTestQueryResult struct {
	QueryString     string
	ExecutionResult string
}

func TestExecute(t *testing.T) {
	t.Parallel()

	log := logger.CreateMock()

	ctx := context.Background()

	e := storage.NewInMemoryEngine()

	st := storage.NewInMemoryStorage(e, log)

	p := compute.NewStdParser()

	cp := compute.NewStdCompute(p, log)

	d := NewDatabase(ctx, st, cp, log)

	tests := map[string]struct {
		testQueryResult []DatabaseTestQueryResult
	}{
		"get set get get delete": {
			testQueryResult: []DatabaseTestQueryResult{
				{
					QueryString:     "GET key1",
					ExecutionResult: "",
				},
				{
					QueryString:     "SET key1 value123",
					ExecutionResult: "1",
				},
				{
					QueryString:     "GET key1",
					ExecutionResult: "value123",
				},
				{
					QueryString:     "GET key2",
					ExecutionResult: "",
				},
				{
					QueryString:     "DEL key1",
					ExecutionResult: "1",
				},
			},
		},
		"delete set get set get delete get": {
			testQueryResult: []DatabaseTestQueryResult{
				{
					QueryString:     "DEL key3",
					ExecutionResult: "1",
				},
				{
					QueryString:     "SET key3 v31",
					ExecutionResult: "1",
				},
				{
					QueryString:     "GET key3",
					ExecutionResult: "v31",
				},
				{
					QueryString:     "SET key3 v42",
					ExecutionResult: "1",
				},
				{
					QueryString:     "GET key3",
					ExecutionResult: "v42",
				},
				{
					QueryString:     "DEL key3",
					ExecutionResult: "1",
				},
				{
					QueryString:     "GET key3",
					ExecutionResult: "",
				},
			},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			for _, dr := range test.testQueryResult {
				res, err := d.Execute(dr.QueryString)

				assert.NoError(t, err)

				assert.Equal(t, res, dr.ExecutionResult)
			}
		})
	}
}
