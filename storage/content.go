package storage

import (
	"context"

	"github.com/libmonsoon-dev/web-crawler/model"
)

type ContentRepository interface {
	Store(context.Context, model.Content) (id int64, err error)
	Load(ctx context.Context, id int64) (model.Content, error)
}
