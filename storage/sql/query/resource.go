package query

import (
	. "github.com/go-jet/jet/v2/postgres"

	"github.com/libmonsoon-dev/web-crawler/model"
	"github.com/libmonsoon-dev/web-crawler/storage/sql"
	. "github.com/libmonsoon-dev/web-crawler/storage/sql/query/postgres/public/table"
)

func (b *PgBuilder) StoreResource(r model.Resource) sql.Statement {
	return Resources.
		INSERT(
			Resources.WebsiteID,
			Resources.Path,
		).
		VALUES(
			r.WebsiteId,
			r.Path,
		).
		RETURNING(
			Resources.ID,
		)
}

func (b *PgBuilder) LoadResource(id int64) sql.Statement {
	return SELECT(
		Resources.ID,
		Resources.WebsiteID,
		Resources.Path,
	).
		FROM(Resources).
		WHERE(Resources.ID.EQ(Int64(id))).
		LIMIT(1)
}
