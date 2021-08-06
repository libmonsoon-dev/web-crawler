package storage

import (
	"context"

	"github.com/libmonsoon-dev/web-crawler/model"
)

type RequestRepository interface {
	Store(context.Context, model.Request) (id int64, err error)
	Load(ctx context.Context, id int64) (model.Request, error)
}
