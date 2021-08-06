package storage

import (
	"context"

	"github.com/libmonsoon-dev/web-crawler/model"
)

type WebsiteRepository interface {
	Store(context.Context, model.Website) (id int64, err error)
	Load(ctx context.Context, id int64) (model.Website, error)
}
