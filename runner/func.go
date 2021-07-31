package runner

import (
	"context"
)

var _ Interface = new(funcRunner)

func Func(fn func(context.Context) error) Interface {
	return funcRunner(fn)
}

type funcRunner func(context.Context) error

func (f funcRunner) Run(ctx context.Context) error {
	return f(ctx)
}
