package useragent

import (
	"sync"
	"time"

	browser "github.com/EDDYCJY/fake-useragent"
)

var (
	sourceOnce sync.Once
	source     userAgentSource
)

type userAgentSource interface {
	Computer() string
}

func initSource() {
	sourceOnce.Do(doInitSource)
}

func doInitSource() {
	client := browser.Client{Timeout: 10 * time.Second}
	cache := browser.Cache{UpdateFile: true}
	source = browser.NewBrowser(client, cache)
}
