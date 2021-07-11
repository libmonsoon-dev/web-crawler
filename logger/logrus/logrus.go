package logrus

import (
	log "github.com/sirupsen/logrus"

	"github.com/libmonsoon-dev/web-crawler/logger"
)

var _ logger.Logger = new(log.Logger)
