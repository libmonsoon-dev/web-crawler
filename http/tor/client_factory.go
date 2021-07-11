package tor

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/cretz/bine/tor"
	"github.com/ipsn/go-libtor"

	"github.com/libmonsoon-dev/web-crawler/logger"
)

type ClientFactory struct {
	t *tor.Tor
	d *tor.Dialer
}

func (cf *ClientFactory) Close() error {
	return cf.t.Close()
}

func (cf *ClientFactory) NewClient() *http.Client {
	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: cf.d.DialContext,
		},
	}

	return httpClient
}

func NewClientFactory(ctx context.Context, logFactory logger.Factory) (*ClientFactory, error) {
	conf := &tor.StartConf{
		ProcessCreator:         libtor.Creator,
		UseEmbeddedControlConn: true,
		TempDataDirBase:        os.TempDir(),
		DebugWriter:            newLogWriter(logFactory, "Tor"),
	}

	t, err := tor.Start(ctx, conf)
	if err != nil {
		return nil, fmt.Errorf("start Tor connection: %w", err)
	}

	d, err := t.Dialer(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("creating dialer: %w", err)
	}

	f := &ClientFactory{t, d}
	return f, nil
}
