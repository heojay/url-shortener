package handler

import (
	"github.com/google/uuid"
	"github.com/heojay/url-shortener/pkg/model"
	"github.com/heojay/url-shortener/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) GetLongUrl(c echo.Context) error {
	req := new(getLongUrlRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	u, err := h.urlStore.GetByShortUrl(req.ShortUrl)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	res := new(getLongUrlResponse)

	if u.LongUrl == "" {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	res.LongUrl = u.LongUrl

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) ShortenUrl(c echo.Context) error {
	req := new(shortenUrlRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	res := new(shortenUrlResponse)

	if u, _ := h.urlStore.GetByLongUrl(req.LongUrl); u != nil {
		res.ShortUrl = u.ShortUrl
		return c.JSON(http.StatusOK, res)
	}

	var u model.Url
	u.ID = uuid.New()
	u.LongUrl = req.LongUrl
	u.ShortUrl = generateShortUrl(u.ID)

	if err := h.urlStore.CreateUrl(&u); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	res.ShortUrl = u.ShortUrl
	return c.JSON(http.StatusOK, res)
}

func generateShortUrl(uuid uuid.UUID) string {
	return utils.EncodeBase62(uuid)
}
