package database

import (
	"balun_homework_1/foundation/compute"
	"balun_homework_1/foundation/logger"
	"context"
	"fmt"
)

type Storage interface {
	Get(ctx context.Context, key string) (string, bool, error)
	Set(ctx context.Context, key string, value string) (bool, error)
	Delete(ctx context.Context, key string) (bool, error)
}

type Database struct {
	ctx context.Context
	st  Storage
	cp  compute.Compute
	log *logger.Logger
}

var (
	DBQueryParseError            = fmt.Errorf("parse query error")
	DBQueryExecutionError        = fmt.Errorf("query execution error")
	DBQueryOperationNotSupported = fmt.Errorf("operation not supported")
)

func NewDatabase(ctx context.Context, st Storage, cp compute.Compute, log *logger.Logger) *Database {
	return &Database{
		ctx: ctx,
		st:  st,
		cp:  cp,
		log: log,
	}
}

func (d *Database) Execute(query string) (string, error) {
	q, err := d.cp.Handle(d.ctx, query)

	d.log.Debug(d.ctx, "[Database::Execute] Try to handle query", query)

	if err != nil {
		return "", DBQueryParseError
	}

	switch q.Cmd {
	case compute.CommandGet:
		v, _, err := d.st.Get(d.ctx, q.Args[0])

		if err != nil {
			return "", DBQueryExecutionError
		}

		return v, nil

	case compute.CommandSet:
		v, err := d.st.Set(d.ctx, q.Args[0], q.Args[1])

		if err != nil {
			return "", DBQueryExecutionError
		}

		if !v {
			return "0", nil
		}

		return "1", nil

	case compute.CommandDel:
		v, err := d.st.Delete(d.ctx, q.Args[0])

		if err != nil {
			return "", DBQueryExecutionError
		}

		if !v {
			return "0", nil
		}

		return "1", nil
	}

	return "", DBQueryOperationNotSupported
}
