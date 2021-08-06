package pgrepo_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/libmonsoon-dev/web-crawler/model"
)

func TestContentRepo(t *testing.T) {
	ctx := context.Background()
	db, err := NewTestDB(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close(ctx)

	content := model.Content{
		Content:   "example.com",
		Type:      "test/html",
		Processed: true,
	}

	content.Id, err = db.contentRepo.Store(ctx, content)
	if err != nil {
		t.Fatal(err)
	}

	got, err := db.contentRepo.Load(ctx, content.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, content) {
		t.Fatalf("Store & Load:\nGot:\t%+v\nWant:\t%+v", got, content)
	}
}

func TestContentRepoStore(t *testing.T) {
	ctx := context.Background()
	db, err := NewTestDB(ctx, t)
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close(ctx)

	content := model.Content{
		Content:   "example.com",
		Type:      "test/html",
		Processed: true,
	}

	for i := 0; i < 3; i++ {
		content.Id, err = db.contentRepo.Store(ctx, content)
		if err != nil {
			t.Fatal(err)
		}
	}

	got, err := db.contentRepo.Load(ctx, content.Id)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, content) {
		t.Fatalf("Store & Load:\nGot:\t%+v\nWant:\t%+v", got, content)
	}
}
