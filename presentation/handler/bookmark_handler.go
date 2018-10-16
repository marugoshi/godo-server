package handler

import (
	"context"
	"github.com/marugoshi/gobm/presentation/httputils"
)

type BookmarkHandler interface {
	Bookmarks(ctx context.Context, http httputils.Http) error
	ShowBookmark(ctx context.Context, http httputils.Http) error
}

type bookmarkHandler struct {
}

func NewBookmarkHandler() BookmarkHandler {
	return &bookmarkHandler{}
}

func (b *bookmarkHandler) Bookmarks(ctx context.Context, http httputils.Http) error {
	http.Text(200, "hoge")
	return nil
}

func (b *bookmarkHandler) ShowBookmark(ctx context.Context, http httputils.Http) error {
	http.Text(200, http.Params[0])
	return nil
}