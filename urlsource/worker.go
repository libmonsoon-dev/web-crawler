package urlsource

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/libmonsoon-dev/web-crawler/logger"
	"github.com/libmonsoon-dev/web-crawler/model"
	"github.com/libmonsoon-dev/web-crawler/storage"
)

func NewWorker(loggerFactory logger.Factory, repo storage.URLRepository) *Worker {
	return &Worker{
		logger: loggerFactory.New("url source"),
		repo:   repo,
		output: make(chan model.URL, 1),

		nothingToSendPause: time.Second * 10,
	}
}

type Worker struct {
	logger logger.Logger
	repo   storage.URLRepository
	output chan model.URL

	nothingToSendPause time.Duration
}

func (w *Worker) GetUrlOutputStream() <-chan model.URL {
	return w.output
}

func (w *Worker) Run(ctx context.Context) error {
	for {
		sent, err := w.run(ctx)
		if err != nil {
			return err
		}

		if sent == 0 {
			time.Sleep(w.nothingToSendPause)
		}
	}
}

func (w *Worker) run(ctx context.Context) (sent int, err error) {
	sent, err = w.repo.GetURLsToVisitStream(ctx, w.getLastVisitTill(), w.output)
	if err != nil && isContextError(err) {
		return sent, fmt.Errorf("send url: %w", err)
	} else if err != nil {
		w.logger.Error("get urls from repo", err)
	}

	return
}

func (w *Worker) getLastVisitTill() time.Time {
	return time.Now() // TODO
}

func isContextError(err error) bool {
	return errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded)
}
