package pgrepo

import (
	"context"

	"github.com/libmonsoon-dev/web-crawler/model"
	"github.com/libmonsoon-dev/web-crawler/storage"
	"github.com/libmonsoon-dev/web-crawler/storage/sql"
)

type ContentQueryBuilder interface {
	StoreContent(website model.Content) sql.Statement
	LoadContent(id int64) sql.Statement
}

func NewContentRepository(db *sql.DB, queryBuilder ContentQueryBuilder) *ContentRepo {
	repo := &ContentRepo{
		common:       common{db},
		db:           db,
		queryBuilder: queryBuilder,
	}
	return repo
}

type ContentRepo struct {
	common
	db           *sql.DB
	queryBuilder ContentQueryBuilder
}

var _ storage.ContentRepository = (*ContentRepo)(nil)

func (c *ContentRepo) Store(ctx context.Context, content model.Content) (id int64, err error) {
	stmt := c.queryBuilder.StoreContent(content)
	return c.store(ctx, stmt)
}

func (c *ContentRepo) Load(ctx context.Context, id int64) (content model.Content, err error) {
	stmt := c.queryBuilder.LoadContent(id)
	err = c.load(ctx, stmt, &content.Id, &content.Content, &content.Type, &content.Processed)

	return
}
