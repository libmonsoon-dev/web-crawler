package pgrepo

import (
	"context"

	"github.com/libmonsoon-dev/web-crawler/model"
	"github.com/libmonsoon-dev/web-crawler/storage"
	"github.com/libmonsoon-dev/web-crawler/storage/sql"
)

type WebsiteQueryBuilder interface {
	StoreWebsite(website model.Website) sql.Statement
	LoadWebsite(id int64) sql.Statement
}

func NewWebsiteRepository(db *sql.DB, queryBuilder WebsiteQueryBuilder) *WebsiteRepository {
	repo := &WebsiteRepository{
		common:       common{db},
		db:           db,
		queryBuilder: queryBuilder,
	}
	return repo
}

var _ storage.WebsiteRepository = (*WebsiteRepository)(nil)

type WebsiteRepository struct {
	common
	db           *sql.DB
	queryBuilder WebsiteQueryBuilder
}

func (w *WebsiteRepository) Store(ctx context.Context, website model.Website) (id int64, err error) {
	stmt := w.queryBuilder.StoreWebsite(website)
	return w.store(ctx, stmt)
}

func (w *WebsiteRepository) Load(ctx context.Context, id int64) (website model.Website, err error) {
	stmt := w.queryBuilder.LoadWebsite(id)
	err = w.load(ctx, stmt, &website.Id, &website.Host)

	return
}
