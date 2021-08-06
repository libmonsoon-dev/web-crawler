package migration

import (
	"embed"
	"fmt"
	"net/http"
	"net/url"

	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
)

const (
	EmbedDriverName = "embed"

	DefaultPath = "."
)

type EmbedDriver struct {
	fs embed.FS
	httpfs.PartialDriver
}

func (e *EmbedDriver) Open(rawURL string) (source.Driver, error) {
	URL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("parse URL: %w", err)
	}

	path := URL.Opaque
	if len(path) == 0 {
		path = URL.Host + URL.Path
	}

	if len(path) == 0 {
		path = DefaultPath
	}

	newDriver := NewEmbedDriver(e.fs)
	if err := newDriver.Init(http.FS(e.fs), path); err != nil {
		return nil, fmt.Errorf("init driver: %w", err)
	}

	return newDriver, nil
}

func NewEmbedDriver(fs embed.FS) *EmbedDriver {
	ed := &EmbedDriver{
		fs: fs,
	}
	return ed
}

var _ source.Driver = (*EmbedDriver)(nil)

//go:embed *.sql
var fs embed.FS

func init() {
	source.Register(EmbedDriverName, NewEmbedDriver(fs))
}
