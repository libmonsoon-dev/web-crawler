package pgrepo_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/libmonsoon-dev/web-crawler/model"
)

func TestRequestRepo(t *testing.T) {
	ctx := context.Background()
	db, err := NewTestDB(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close(ctx)

	date := time.Date(2009, 11, 17, 20, 34, 58, 0, time.Local)
	request := model.Request{
		Started:    date,
		Ended:      date.Add(7 * time.Second),
		Headers:    http.Header{"Content-Type": []string{"text/html"}},
		StatusCode: 200,
	}

	request.WebsiteId, err = db.websiteRepo.Store(ctx, model.Website{Host: "example.com"})
	if err != nil {
		t.Fatal(err)
	}

	resource := model.Resource{
		WebsiteId: request.WebsiteId,
		Path:      "/",
	}
	request.ResourceId, err = db.resourceRepo.Store(ctx, resource)
	if err != nil {
		t.Fatal(err)
	}

	content := model.Content{
		Content:   "example.com",
		Type:      "test/html",
		Processed: true,
	}
	request.ContentId, err = db.contentRepo.Store(ctx, content)
	if err != nil {
		t.Fatal(err)
	}

	request.Id, err = db.requestRepo.Store(ctx, request)
	if err != nil {
		t.Fatal(err)
	}

	got, err := db.requestRepo.Load(ctx, request.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !requestsEqual(got, request) {
		t.Fatalf("Store & Load:\nGot:\t%+v\nWant:\t%+v", got, request)
	}
}

func requestsEqual(a model.Request, b model.Request) bool {
	return a.Id == b.Id &&
		a.WebsiteId == b.WebsiteId &&
		a.ResourceId == b.ResourceId &&
		a.ContentId == b.ContentId &&
		datesEqual(a.Started, b.Started) &&
		datesEqual(a.Ended, b.Ended) &&
		reflect.DeepEqual(a.Headers, b.Headers) &&
		a.StatusCode == b.StatusCode
}

func datesEqual(a time.Time, b time.Time) bool {
	return a.Unix() == b.Unix()
}
