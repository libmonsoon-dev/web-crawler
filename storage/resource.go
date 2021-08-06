package storage

import (
	"context"

	"github.com/libmonsoon-dev/web-crawler/model"
)

type ResourceRepository interface {
	Store(context.Context, model.Resource) (id int64, err error)
	Load(ctx context.Context, id int64) (model.Resource, error)
}
