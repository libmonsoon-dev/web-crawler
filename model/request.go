package model

import (
	"net/http"
	"time"
)

type Request struct {
	Id         int64
	WebsiteId  int64
	ResourceId int64
	ContentId  int64
	Started    time.Time
	Ended      time.Time
	Headers    http.Header
	StatusCode int
}
