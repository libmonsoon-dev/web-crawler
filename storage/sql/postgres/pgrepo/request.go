package pgrepo

import (
	"context"

	"github.com/lib/pq/hstore"

	"github.com/libmonsoon-dev/web-crawler/model"
	"github.com/libmonsoon-dev/web-crawler/storage"
	"github.com/libmonsoon-dev/web-crawler/storage/sql"
	"github.com/libmonsoon-dev/web-crawler/storage/sql/postgres/pgutils"
)

type RequestQueryBuilder interface {
	StoreRequest(request model.Request) sql.Statement
	LoadRequest(id int64) sql.Statement
}

func NewRequestRepository(db *sql.DB, queryBuilder RequestQueryBuilder) *RequestRepo {
	repo := &RequestRepo{
		common:       common{db},
		db:           db,
		queryBuilder: queryBuilder,
	}
	return repo
}

type RequestRepo struct {
	common
	db           *sql.DB
	queryBuilder RequestQueryBuilder
}

var _ storage.RequestRepository = (*RequestRepo)(nil)

func (r *RequestRepo) Store(ctx context.Context, request model.Request) (id int64, err error) {
	stmt := r.queryBuilder.StoreRequest(request)
	return r.store(ctx, stmt)
}

func (r *RequestRepo) Load(ctx context.Context, id int64) (req model.Request, err error) {
	stmt := r.queryBuilder.LoadRequest(id)

	headers := hstore.Hstore{}
	err = r.load(
		ctx,
		stmt,
		&req.Id, &req.WebsiteId, &req.ResourceId, &req.ContentId, &req.Started, &req.Ended,
		&headers, &req.StatusCode,
	)

	req.Headers = pgutils.HstoreToHeaders(headers)

	return
}
