package storage

import (
	"context"
	"time"

	"github.com/libmonsoon-dev/web-crawler/model"
)

type URLRepository interface {
	GetURLsToVisitStream(ctx context.Context, lastVisitNotLater time.Time, output chan<- model.URL) (n int, err error)
}
