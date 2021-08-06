package sql

import (
	"database/sql"
)

type DB = sql.DB
type ColumnType = sql.ColumnType

var ErrNoRows = sql.ErrNoRows
