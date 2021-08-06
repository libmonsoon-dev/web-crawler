package query

import (
	"github.com/go-jet/jet/v2/postgres"
)

func asIs(c postgres.Column) postgres.Projection {
	return c.AS(c.Name())
}
