package handler

import (
	"github.com/heojay/url-shortener/pkg/url"
)

type Handler struct {
	urlStore url.Store
}

func New(us url.Store) *Handler {
	return &Handler{
		urlStore: us,
	}
}
