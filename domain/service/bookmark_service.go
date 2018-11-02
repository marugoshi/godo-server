package service

import (
	"context"
	"github.com/marugoshi/gobm/domain/model"
)

type BookmarkService interface {
	Bookmarks(ctx context.Context) (interface{}, error)
	Bookmark(ctx context.Context, id int) (interface{}, error)
}

type bookmarkService struct {
	model.BookmarkModel
}

func NewBookmarkService(m model.BookmarkModel) BookmarkService {
	return &bookmarkService{m}
}

func (b *bookmarkService) Bookmarks(ctx context.Context) (interface{}, error) {
	bookmarks, err := b.BookmarkModel.All(ctx, 1, 100)
	if err != nil {
		return nil, nil
	}
	return bookmarks, nil
}

func (b *bookmarkService) Bookmark(ctx context.Context, id int) (interface{}, error) {
	data := struct {
		Key string
	}{
		Key: "hoge",
	}
	return data, nil
}
