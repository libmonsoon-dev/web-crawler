package runner

import (
	"context"

	"golang.org/x/sync/errgroup"
)

var _ Interface = new(Gather)

type Gather []Interface

func (g Gather) Run(ctx context.Context) error {
	var group *errgroup.Group

	group, ctx = errgroup.WithContext(ctx)
	for _, job := range g {
		group.Go(wrap(ctx, job))
	}

	return group.Wait()
}

func wrap(ctx context.Context, job Interface) func() error {
	return func() error { return job.Run(ctx) }
}
