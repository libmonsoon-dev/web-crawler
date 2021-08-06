package pgrepo

import (
	"context"
	"fmt"
	"time"

	"github.com/libmonsoon-dev/web-crawler/model"
	"github.com/libmonsoon-dev/web-crawler/storage"
	"github.com/libmonsoon-dev/web-crawler/storage/sql"
)

type URLQueryBuilder interface {
	GetURLsToVisit(later time.Time) sql.Statement
}

func NewURLRepository(db *sql.DB, queryBuilder URLQueryBuilder) *URLRepository {
	repo := &URLRepository{
		common:       common{db},
		db:           db,
		queryBuilder: queryBuilder,
	}
	return repo
}

var _ storage.URLRepository = (*URLRepository)(nil)

type URLRepository struct {
	common
	db           *sql.DB
	queryBuilder URLQueryBuilder
}

func (repo *URLRepository) GetURLsToVisitStream(ctx context.Context, lastVisitNotLater time.Time, output chan<- model.URL) (n int, err error) {
	stmt := repo.queryBuilder.GetURLsToVisit(lastVisitNotLater)

	var url model.URL
	scanAndSend := func(r row) (err error) {
		err = r.Scan(&url.WebsiteId, &url.ResourceId, &url.URL)
		if err != nil {
			return fmt.Errorf("scan %s: %w", stmt.DebugSql(), err)
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case output <- url:
			return nil
		}
	}
	return repo.loadMany(ctx, stmt, scanAndSend)
}
