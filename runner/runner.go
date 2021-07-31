package runner

import (
	"context"
)

type Interface interface {
	Run(ctx context.Context) error
}
