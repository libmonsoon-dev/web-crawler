package query

import (
	. "github.com/go-jet/jet/v2/postgres"

	"github.com/libmonsoon-dev/web-crawler/model"
	"github.com/libmonsoon-dev/web-crawler/storage/sql"
	. "github.com/libmonsoon-dev/web-crawler/storage/sql/query/postgres/public/table"
)

func (b *PgBuilder) StoreWebsite(ws model.Website) sql.Statement {
	return Websites.
		INSERT(
			Websites.Host,
		).
		VALUES(
			ws.Host,
		).
		RETURNING(
			Websites.ID,
		)
}

func (b *PgBuilder) LoadWebsite(id int64) sql.Statement {
	return SELECT(
		Websites.ID,
		Websites.Host,
	).
		FROM(Websites).
		WHERE(Websites.ID.EQ(Int64(id))).
		LIMIT(1)
}
