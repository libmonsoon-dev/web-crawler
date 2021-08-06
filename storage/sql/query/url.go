package query

import (
	"time"

	. "github.com/go-jet/jet/v2/postgres"

	"github.com/libmonsoon-dev/web-crawler/storage/sql"
	. "github.com/libmonsoon-dev/web-crawler/storage/sql/query/postgres/public/table"
	. "github.com/libmonsoon-dev/web-crawler/storage/sql/query/postgres/public/view"
)

const maxStartedField = "max_started"

func (b *PgBuilder) GetURLsToVisit(lastVisitNotLater time.Time) sql.Statement {
	maxStartedField := RawTimestampz(maxStartedField)
	return SELECT(
		Urls.WebsiteID,
		Urls.ResourceID,
		Urls.URL,
	).FROM(
		b.urlsWithLatestRequest().AsTable(Urls.TableName()),
	).WHERE(
		maxStartedField.LT(TimestampzT(lastVisitNotLater)).
			OR(maxStartedField.IS_NULL()),
	).ORDER_BY(
		maxStartedField.IS_NOT_NULL(),
		maxStartedField,
	)
}

func (b *PgBuilder) urlsWithLatestRequest() SelectStatement {
	return SELECT(
		asIs(Urls.WebsiteID),
		asIs(Urls.ResourceID),
		asIs(Urls.URL),
		MAX(Requests.Started).AS(maxStartedField),
	).FROM(
		Urls.LEFT_JOIN(Requests, Requests.
			WebsiteID.EQ(Urls.WebsiteID).
			AND(Requests.ResourceID.EQ(Urls.ResourceID))),
	).GROUP_BY(
		Urls.WebsiteID,
		Urls.ResourceID,
		Urls.URL,
	)
}
