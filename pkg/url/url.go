package url

import "github.com/heojay/url-shortener/pkg/model"

type Store interface {
	GetByShortUrlPath(url string) (*model.Url, error)
	GetByOrgUrl(url string) (*model.Url, error)
	CreateUrl(u *model.Url) error
}
