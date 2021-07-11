package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/libmonsoon-dev/web-crawler/crawler"
	"github.com/libmonsoon-dev/web-crawler/http/tor"
	"github.com/libmonsoon-dev/web-crawler/logger"
	"github.com/libmonsoon-dev/web-crawler/logger/logrus"
	"github.com/libmonsoon-dev/web-crawler/run"
)

func main() {
	logFactory := logrus.NewFactory()
	err := Main(logFactory)
	if err != nil {
		logFactory.New("main").Error(err)
		os.Exit(1)
	}
}

func Main(logFactory logger.Factory) error {
	// Used SIGTSTP because all usual signals (SIGINT, SIGTERM, SIGHUP) already handled by Tor
	ctx, stopNotify := signal.NotifyContext(context.Background(), syscall.SIGTSTP)
	defer stopNotify()

	clientFactory, err := tor.NewClientFactory(ctx, logFactory)
	if err != nil {
		return fmt.Errorf("new http client factory: %w", err)
	}
	defer clientFactory.Close()

	c := crawler.NewCrawler(
		logFactory,
		clientFactory,
	)
	return run.Errorf("crawler", c).Run(ctx)
}
