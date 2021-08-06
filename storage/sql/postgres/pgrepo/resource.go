package pgrepo

import (
	"context"

	"github.com/libmonsoon-dev/web-crawler/model"
	"github.com/libmonsoon-dev/web-crawler/storage"
	"github.com/libmonsoon-dev/web-crawler/storage/sql"
)

type ResourceQueryBuilder interface {
	StoreResource(resource model.Resource) sql.Statement
	LoadResource(id int64) sql.Statement
}

func NewResourceRepository(db *sql.DB, queryBuilder ResourceQueryBuilder) *ResourceRepo {
	repo := &ResourceRepo{
		common:       common{db},
		db:           db,
		queryBuilder: queryBuilder,
	}
	return repo
}

var _ storage.ResourceRepository = (*ResourceRepo)(nil)

type ResourceRepo struct {
	common
	db           *sql.DB
	queryBuilder ResourceQueryBuilder
}

func (r *ResourceRepo) Store(ctx context.Context, resource model.Resource) (id int64, err error) {
	stmt := r.queryBuilder.StoreResource(resource)
	return r.store(ctx, stmt)
}

func (r *ResourceRepo) Load(ctx context.Context, id int64) (resource model.Resource, err error) {
	stmt := r.queryBuilder.LoadResource(id)
	err = r.load(ctx, stmt, &resource.Id, &resource.WebsiteId, &resource.Path)

	return
}
