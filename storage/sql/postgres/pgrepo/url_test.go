package pgrepo_test

import (
	"context"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/libmonsoon-dev/web-crawler/model"
)

func TestURLRepository_GetURLsToVisitStream(t *testing.T) {
	ctx := context.Background()

	db, err := NewTestDB(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close(ctx)

	lastVisitNotLater := time.Date(2021, 1, 15, 3, 4, 5, 0, time.Local)
	output := make(chan model.URL)

	site := model.Website{Host: "example.com"}
	site.Id, err = db.websiteRepo.Store(ctx, site)
	if err != nil {
		t.Fatal(err)
	}

	google := model.Website{Host: "google.com"}
	google.Id, err = db.websiteRepo.Store(ctx, google)
	if err != nil {
		t.Fatal(err)
	}

	resourceA := model.Resource{WebsiteId: site.Id, Path: "/"}
	resourceA.Id, err = db.resourceRepo.Store(ctx, resourceA)
	if err != nil {
		t.Fatal(err)
	}

	resourceB := model.Resource{WebsiteId: site.Id, Path: "/index"}
	resourceB.Id, err = db.resourceRepo.Store(ctx, resourceB)
	if err != nil {
		t.Fatal(err)
	}

	resourceC := model.Resource{WebsiteId: site.Id, Path: "/page"}
	resourceC.Id, err = db.resourceRepo.Store(ctx, resourceC)
	if err != nil {
		t.Fatal(err)
	}

	resourceD := model.Resource{WebsiteId: google.Id, Path: "/"}
	resourceD.Id, err = db.resourceRepo.Store(ctx, resourceD)
	if err != nil {
		t.Fatal(err)
	}

	contentA := model.Content{
		Content:   "content",
		Type:      "text/plain",
		Processed: true,
	}
	contentA.Id, err = db.contentRepo.Store(ctx, contentA)
	if err != nil {
		t.Fatal(err)
	}

	contentB := model.Content{
		Content: `<!DOCTYPE html><html><body><h1>h1</h1><p>p</p></body></html>`,
		Type:    "text/html",
	}
	contentB.Id, err = db.contentRepo.Store(ctx, contentB)
	if err != nil {
		t.Fatal(err)
	}

	contentD := model.Content{
		Content: `<!DOCTYPE html><html><body><h1>h1</h1></body></html>`,
		Type:    "text/html",
	}
	contentD.Id, err = db.contentRepo.Store(ctx, contentD)
	if err != nil {
		t.Fatal(err)
	}

	requestA := model.Request{
		WebsiteId:  site.Id,
		ResourceId: resourceA.Id,
		ContentId:  contentA.Id,
		Started:    lastVisitNotLater.Add(time.Second),
		Ended:      lastVisitNotLater.Add(2 * time.Second),
		Headers:    http.Header{"Content-Type": []string{contentA.Type}},
		StatusCode: 200,
	}
	requestA.Id, err = db.requestRepo.Store(ctx, requestA)

	requestB := model.Request{
		WebsiteId:  site.Id,
		ResourceId: resourceB.Id,
		ContentId:  contentB.Id,
		Started:    lastVisitNotLater.Add(-3 * time.Second),
		Ended:      lastVisitNotLater.Add(time.Second),
		Headers:    http.Header{"Content-Type": []string{contentB.Type}},
		StatusCode: 200,
	}
	requestB.Id, err = db.requestRepo.Store(ctx, requestB)

	requestD := model.Request{
		WebsiteId:  google.Id,
		ResourceId: resourceD.Id,
		ContentId:  contentD.Id,
		Started:    lastVisitNotLater.Add(-time.Second),
		Ended:      lastVisitNotLater,
		Headers:    http.Header{"Content-Type": []string{contentD.Type}},
		StatusCode: 200,
	}
	requestD.Id, err = db.requestRepo.Store(ctx, requestD)

	urlsChan := make(chan []model.URL)
	go urlStreamReadAll(output, urlsChan)

	gotN, err := db.urlRepo.GetURLsToVisitStream(ctx, lastVisitNotLater, output)
	close(output)
	if err != nil {
		t.Fatalf("GetURLsToVisitStream() error = %v", err)
	}

	wantN := 3
	if gotN != wantN {
		t.Errorf("GetURLsToVisitStream() gotN = %v, want %v", gotN, wantN)
	}

	var gotURLs []model.URL
	select {
	case <-time.After(time.Second):
		t.Log("wait urls timeout")
	case gotURLs = <-urlsChan:
		// continue
	}

	expectURLs := []model.URL{
		// First returns URL without request
		{URL: site.Host + resourceC.Path, WebsiteId: site.Id, ResourceId: resourceC.Id},
		// then we return the URL with the earliest last request
		{URL: site.Host + resourceB.Path, WebsiteId: site.Id, ResourceId: resourceB.Id},
		{URL: google.Host + resourceD.Path, WebsiteId: google.Id, ResourceId: resourceD.Id},
	}

	if !reflect.DeepEqual(gotURLs, expectURLs) {
		t.Fatalf("GetURLsToVisit:\nGot:\t%+v\nWant:\t%+v", gotURLs, expectURLs)
	}
}

func urlStreamReadAll(input <-chan model.URL, output chan<- []model.URL) {
	defer close(output)
	var buf []model.URL

	for url := range input {
		buf = append(buf, url)
	}

	output <- buf
}
