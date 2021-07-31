package crawler

import (
	"context"

	"github.com/libmonsoon-dev/web-crawler/http"
	"github.com/libmonsoon-dev/web-crawler/logger"
	"github.com/libmonsoon-dev/web-crawler/runner"
)

type Crawler struct {
	logger            logger.Logger
	httpClientFactory http.ClientFactory
	urlSource         UrlSource
}

func (c *Crawler) Run(ctx context.Context) error {
	c.logger.Trace("start", ctx)
	defer c.logger.Trace("exit", ctx)

	g := runner.Gather{
		runner.Errorf("URL source", c.urlSource),
	}
	return g.Run(ctx)
}

func NewCrawler(logFactory logger.Factory, httpClientFactory http.ClientFactory, urlSource UrlSource) *Crawler {
	c := &Crawler{
		logger:            logFactory.New("crawler"),
		httpClientFactory: httpClientFactory,
		urlSource:         urlSource,
	}
	return c
}
