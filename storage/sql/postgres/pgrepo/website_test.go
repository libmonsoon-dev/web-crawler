package pgrepo_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/libmonsoon-dev/web-crawler/model"
)

func TestNewWebsiteRepository(t *testing.T) {
	ctx := context.Background()
	db, err := NewTestDB(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close(ctx)

	site := model.Website{
		Host: "example.com",
	}

	site.Id, err = db.websiteRepo.Store(ctx, site)
	if err != nil {
		t.Fatal(err)
	}

	got, err := db.websiteRepo.Load(ctx, site.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, site) {
		t.Fatalf("Store & Load:\nGot:\t%+v\nWant:\t%+v", got, site)
	}

}
