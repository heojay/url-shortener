package handler

import (
	"github.com/google/uuid"
	"github.com/heojay/url-shortener/pkg/model"
	"github.com/heojay/url-shortener/pkg/utils"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *Handler) GetOrgUrl(c echo.Context) error {
	req := new(getOrgUrlRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	u, err := h.urlStore.GetByShortUrlPath(req.ShortUrlPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	res := new(getOrgUrlResponse)

	if u.OrgUrl == "" {
		return c.JSON(http.StatusNotFound, utils.NotFound())
	}
	res.OrgUrl = u.OrgUrl

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) ShortenUrl(c echo.Context) error {
	req := new(shortenUrlRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	res := new(shortenUrlResponse)

	if u, _ := h.urlStore.GetByOrgUrl(req.OrgUrl); u != nil {
		res.ShortUrlPath = u.ShortUrlPath
		return c.JSON(http.StatusOK, res)
	}

	var u model.Url
	u.Uuid = uuid.New()
	u.OrgUrl = req.OrgUrl
	u.ShortUrlPath = generateShortUrlPath(u.Uuid)

	if err := h.urlStore.CreateUrl(&u); err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}

	res.ShortUrlPath = u.ShortUrlPath
	return c.JSON(http.StatusOK, res)
}

func generateShortUrlPath(uuid uuid.UUID) string {
	return utils.EncodeBase62(uuid)
}
