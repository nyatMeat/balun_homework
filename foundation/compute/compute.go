package compute

import (
	"balun_homework_1/foundation/logger"
	"context"
	"fmt"
)

type Compute interface {
	Handle(ctx context.Context, query string) (Query, error)
}

var (
	ParseStringError = fmt.Errorf("parse error")
)

type StdCompute struct {
	p   *StdParser
	log *logger.Logger
}

func NewStdCompute(p *StdParser, log *logger.Logger) *StdCompute {
	return &StdCompute{
		p:   p,
		log: log,
	}
}

func (c *StdCompute) Handle(ctx context.Context, query string) (Query, error) {
	c.log.Debug(ctx, "[StdCompute::Execute] Try to Parse query string", query)

	q, err := c.p.Parse(query)

	if err != nil {
		c.log.Warn(ctx, "[StdCompute::Execute] Error during query string parsing", query, err)

		return Query{}, ParseStringError
	}

	return q, nil
}
