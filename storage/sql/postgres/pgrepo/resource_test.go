package pgrepo_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/libmonsoon-dev/web-crawler/model"
)

func TestResourceRepo(t *testing.T) {
	ctx := context.Background()
	db, err := NewTestDB(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close(ctx)

	resource := model.Resource{
		Path: "/",
	}

	resource.WebsiteId, err = db.websiteRepo.Store(ctx, model.Website{Host: "example.com"})
	if err != nil {
		t.Fatal(err)
	}

	resource.Id, err = db.resourceRepo.Store(ctx, resource)
	if err != nil {
		t.Fatal(err)
	}

	got, err := db.resourceRepo.Load(ctx, resource.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, resource) {
		t.Fatalf("Store & Load:\nGot:\t%+v\nWant:\t%+v", got, resource)
	}
}
