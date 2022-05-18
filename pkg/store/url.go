package store

import (
	"github.com/heojay/url-shortener/pkg/model"
	"gorm.io/gorm"
)

type UrlStore struct {
	db *gorm.DB
}

func NewUrlStore(db *gorm.DB) *UrlStore {
	return &UrlStore{
		db: db,
	}
}

func (us *UrlStore) GetByShortUrl(url string) (*model.Url, error) {
	var m model.Url
	if err := us.db.Where(&model.Url{ShortUrl: url}).First(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UrlStore) GetByLongUrl(url string) (*model.Url, error) {
	var m model.Url
	if err := us.db.Where(&model.Url{LongUrl: url}).First(&m).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (us *UrlStore) CreateUrl(u *model.Url) error {
	return us.db.Create(&u).Error
}
