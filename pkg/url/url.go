package url

import "github.com/heojay/url-shortener/pkg/model"

type Store interface {
	GetByShortUrl(url string) (*model.Url, error)
	GetByLongUrl(url string) (*model.Url, error)
	CreateUrl(u *model.Url) error
}
