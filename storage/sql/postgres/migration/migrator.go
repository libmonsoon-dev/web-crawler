package migration

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"sync"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/libmonsoon-dev/web-crawler/storage"
)

func NewMigrator(db *sql.DB) *Migrator {
	m := &Migrator{
		db: db,
	}
	return m
}

type Migrator struct {
	db *sql.DB

	initOnce sync.Once
	initErr  error

	migrator *migrate.Migrate
}

var _ storage.Migrator = (*Migrator)(nil)

func (m *Migrator) Up(ctx context.Context) error {
	err := m.init()
	if err != nil {
		return fmt.Errorf("init migrator: %w", err)
	}

	return m.call(ctx, m.migrator.Up)
}

func (m *Migrator) Down(ctx context.Context) error {
	err := m.init()
	if err != nil {
		return fmt.Errorf("init migrator: %w", err)
	}

	return m.call(ctx, m.migrator.Down)
}

func (m *Migrator) init() error {
	m.initOnce.Do(func() {
		err := m.doInit()
		if err != nil {
			m.initErr = fmt.Errorf("init: %w", err)
		}
	})

	return m.initErr
}

func (m *Migrator) Close() error {
	_, dbErr := m.migrator.Close()
	return dbErr
}

func (m *Migrator) doInit() error {
	driver, err := postgres.WithInstance(m.db, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("new postgres instance: %w", err)
	}

	migrationSource := &url.URL{
		Scheme: EmbedDriverName,
		Path:   DefaultPath,
	}

	m.migrator, err = migrate.NewWithDatabaseInstance(
		migrationSource.String(),
		"",
		driver,
	)
	if err != nil {
		return fmt.Errorf("new postgres migrator: %w", err)
	}

	return nil
}

func (m *Migrator) call(ctx context.Context, fn func() error) error {
	innerCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		<-innerCtx.Done()

		select {
		case <-ctx.Done():
			m.migrator.GracefulStop <- true
		default:
			// closed internal context on method exit
		}
	}()

	err := fn()
	if err != nil {
		return fmt.Errorf("migrator call: %w", err)
	}

	return ctx.Err()
}
