package crawler

import (
	"context"

	"github.com/libmonsoon-dev/web-crawler/http"
	"github.com/libmonsoon-dev/web-crawler/logger"
)

type Crawler struct {
	logger            logger.Logger
	httpClientFactory http.ClientFactory
}

func (c Crawler) Run(ctx context.Context) error {
	c.logger.Trace("start", ctx)
	<-ctx.Done()
	c.logger.Trace("exit", ctx)

	return ctx.Err()
}

func NewCrawler(logFactory logger.Factory, httpClientFactory http.ClientFactory) *Crawler {
	c := &Crawler{
		logger:            logFactory.New("crawler"),
		httpClientFactory: httpClientFactory,
	}
	return c
}
