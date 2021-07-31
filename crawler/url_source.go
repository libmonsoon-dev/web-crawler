package crawler

import (
	"github.com/libmonsoon-dev/web-crawler/model"
	"github.com/libmonsoon-dev/web-crawler/runner"
)

type UrlSource interface {
	GetUrlOutputStream() <-chan model.URL
	runner.Interface
}
