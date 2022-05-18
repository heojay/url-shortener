package main

import (
	"github.com/heojay/url-shortener/pkg/db"
	"github.com/heojay/url-shortener/pkg/handler"
	"github.com/heojay/url-shortener/pkg/model"
	"github.com/heojay/url-shortener/pkg/router"
	"github.com/heojay/url-shortener/pkg/store"
	"github.com/labstack/gommon/log"
)

func main() {
	r := router.New()

	v1 := r.Group("/api")

	d, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	err = d.AutoMigrate(&model.Url{})
	if err != nil {
		log.Fatal(err)
	}

	us := store.NewUrlStore(d)
	h := handler.New(us)
	h.Register(v1)

	r.Logger.Fatal(r.Start(":8000"))
}
