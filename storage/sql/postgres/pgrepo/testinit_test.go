package pgrepo_test

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"net/url"
	"os"
	"testing"
	"time"

	"github.com/libmonsoon-dev/web-crawler/storage"
	"github.com/libmonsoon-dev/web-crawler/storage/sql/postgres/migration"
	"github.com/libmonsoon-dev/web-crawler/storage/sql/postgres/pgrepo"
	"github.com/libmonsoon-dev/web-crawler/storage/sql/query"
)

func NewTestDB(ctx context.Context, tb testing.TB) (*DB, error) {
	const connEnvKey = "TEST_PG_CONNECTION"

	dataSource := os.Getenv(connEnvKey)
	if dataSource == "" {
		tb.Skipf("Environment variable %q is empty, skipping", connEnvKey)
	}

	dbName := "test_" + randomString(12)
	adminDB, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, fmt.Errorf("open %v: %w", dataSource, err)
	}

	_, err = adminDB.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE %v WITH ENCODING 'utf-8'", dbName))
	if err != nil {
		return nil, fmt.Errorf("create test db: %w", err)
	}

	URL, err := url.Parse(dataSource)
	if err != nil {
		return nil, fmt.Errorf("parse data source url: %w", err)
	}

	URL.Path = dbName
	dataSource = URL.String()

	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		return nil, fmt.Errorf("open %v: %w", dataSource, err)
	}

	migrator := migration.NewMigrator(db)
	err = migrator.Up(ctx)
	if err != nil {
		return nil, fmt.Errorf("apply migrations: %w", err)
	}

	queryBuilder := query.NewPGBuilder()

	r := &DB{
		tb:           tb,
		adminDB:      adminDB,
		name:         dbName,
		db:           db,
		migrator:     migrator,
		websiteRepo:  pgrepo.NewWebsiteRepository(db, queryBuilder),
		resourceRepo: pgrepo.NewResourceRepository(db, queryBuilder),
		contentRepo:  pgrepo.NewContentRepository(db, queryBuilder),
		requestRepo:  pgrepo.NewRequestRepository(db, queryBuilder),
		urlRepo:      pgrepo.NewURLRepository(db, queryBuilder),
	}
	return r, nil
}

type DB struct {
	tb           testing.TB
	adminDB      *sql.DB
	name         string
	db           *sql.DB
	migrator     storage.Migrator
	websiteRepo  storage.WebsiteRepository
	resourceRepo storage.ResourceRepository
	contentRepo  storage.ContentRepository
	requestRepo  storage.RequestRepository
	urlRepo      storage.URLRepository
}

func (db *DB) Close(ctx context.Context) {
	err := db.migrator.Down(ctx)
	if err != nil {
		err = fmt.Errorf("rollback migrations: %w", err)
		db.tb.Error(err)
	}

	err = db.migrator.Close()
	if err != nil {
		err = fmt.Errorf("close migrator: %w", err)
		db.tb.Error(err)
	}

	err = db.db.Close()
	if err != nil {
		err = fmt.Errorf("db close: %w", err)
		db.tb.Error(err)
	}

	_, err = db.adminDB.ExecContext(ctx, fmt.Sprintln("DROP DATABASE", db.name))
	if err != nil {
		err = fmt.Errorf("drop test db: %w", err)
		db.tb.Error(err)
	}

	err = db.adminDB.Close()
	if err != nil {
		err = fmt.Errorf("admin db close: %w", err)
		db.tb.Error(err)
	}

	return
}

var r = rand.New(rand.NewSource(time.Now().Unix()))

func randomString(length int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyz")

	b := make([]rune, length)
	for i := range b {
		b[i] = letterRunes[r.Intn(len(letterRunes))]
	}
	return string(b)
}
