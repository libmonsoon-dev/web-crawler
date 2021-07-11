package tor

import (
	"github.com/libmonsoon-dev/web-crawler/logger"
)

func newLogWriter(logFactory logger.Factory, component string) logWriter {
	l := logFactory.New(component)
	return logWriter{l}
}

type logWriter struct {
	logger.Logger
}

func (l logWriter) Write(p []byte) (n int, err error) {
	l.Debug(string(p))
	return len(p), err
}
