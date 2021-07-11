package logrus

import (
	log "github.com/sirupsen/logrus"

	"github.com/libmonsoon-dev/web-crawler/logger"
)

type factory struct {
	l *log.Logger
}

func (f factory) New(component string) logger.Logger {
	return f.l.WithField("component", component)
}

func NewFactory() logger.Factory {
	l := log.New()
	l.SetLevel(log.DebugLevel)
	l.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	return factory{l}
}
