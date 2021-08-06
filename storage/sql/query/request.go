package query

import (
	. "github.com/go-jet/jet/v2/postgres"

	"github.com/libmonsoon-dev/web-crawler/model"
	"github.com/libmonsoon-dev/web-crawler/storage/sql"
	"github.com/libmonsoon-dev/web-crawler/storage/sql/postgres/pgutils"
	. "github.com/libmonsoon-dev/web-crawler/storage/sql/query/postgres/public/table"
)

func (b *PgBuilder) StoreRequest(r model.Request) sql.Statement {
	return Requests.INSERT(
		Requests.WebsiteID,
		Requests.ResourceID,
		Requests.ContentID,
		Requests.Started,
		Requests.Ended,
		Requests.Headers,
		Requests.StatusCode,
	).VALUES(
		r.WebsiteId,
		r.ResourceId,
		r.ContentId,
		r.Started,
		r.Ended,
		pgutils.HeadersToHstore(r.Headers),
		r.StatusCode,
	).RETURNING(
		Requests.ID,
	)
}

func (b *PgBuilder) LoadRequest(id int64) sql.Statement {
	return SELECT(
		Requests.ID,
		Requests.WebsiteID,
		Requests.ResourceID,
		Requests.ContentID,
		Requests.Started,
		Requests.Ended,
		Requests.Headers,
		Requests.StatusCode,
	).
		FROM(Requests).
		WHERE(Requests.ID.EQ(Int64(id))).
		LIMIT(1)
}
