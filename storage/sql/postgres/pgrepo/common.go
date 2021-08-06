package pgrepo

import (
	"context"
	"errors"
	"fmt"

	"github.com/libmonsoon-dev/web-crawler/storage"
	"github.com/libmonsoon-dev/web-crawler/storage/sql"
)

type common struct {
	db *sql.DB
}

func (c common) store(ctx context.Context, stmt sql.Statement) (id int64, err error) {
	query, args := stmt.Sql()

	row := c.db.QueryRowContext(ctx, query, args...)

	err = row.Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return
}

func (c common) load(ctx context.Context, stmt sql.Statement, dest ...interface{}) error {
	query, args := stmt.Sql()
	row := c.db.QueryRowContext(ctx, query, args...)

	err := row.Scan(dest...)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return storage.NoData
	}
	if err != nil {
		return fmt.Errorf("scan %s result: %w", stmt.DebugSql(), err)
	}

	return nil
}

func (c common) loadMany(ctx context.Context, stmt sql.Statement, fn func(row) error) (n int, err error) {
	query, args := stmt.Sql()
	rows, err := c.db.QueryContext(ctx, query, args...)
	if err != nil {
		return n, fmt.Errorf("query %s: %w", stmt.DebugSql(), err)
	}
	defer rows.Close()

	for rows.Next() {
		err = fn(rows)
		if err != nil {
			return n, fmt.Errorf("callback: %w", err)
		}
		n++
	}
	return n, rows.Err()
}

type row interface {
	Scan(dest ...interface{}) error
	ColumnTypes() ([]*sql.ColumnType, error)
	Columns() ([]string, error)
}
