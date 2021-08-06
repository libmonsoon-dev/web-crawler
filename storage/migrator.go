package storage

import (
	"context"
	"io"
)

type Migrator interface {
	Up(context.Context) error
	Down(context.Context) error

	io.Closer
}
