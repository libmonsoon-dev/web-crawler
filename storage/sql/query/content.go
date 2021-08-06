package query

import (
	. "github.com/go-jet/jet/v2/postgres"

	"github.com/libmonsoon-dev/web-crawler/model"
	"github.com/libmonsoon-dev/web-crawler/storage/sql"
	. "github.com/libmonsoon-dev/web-crawler/storage/sql/query/postgres/public/table"
)

func (b *PgBuilder) StoreContent(c model.Content) sql.Statement {
	return Contents.
		INSERT(
			Contents.Content,
			Contents.Type,
			Contents.Processed,
		).
		VALUES(
			c.Content,
			c.Type,
			c.Processed,
		).
		ON_CONFLICT(Contents.Content).
		DO_UPDATE(
			SET(Contents.Type.SET(String(c.Type))),
		).
		RETURNING(
			Contents.ID,
		)
}

func (b *PgBuilder) LoadContent(id int64) sql.Statement {
	return SELECT(
		Contents.ID,
		Contents.Content,
		Contents.Type,
		Contents.Processed,
	).
		FROM(Contents).
		WHERE(Contents.ID.EQ(Int64(id))).
		LIMIT(1)
}
