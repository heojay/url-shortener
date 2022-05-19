package handler

import "github.com/labstack/echo/v4"

func (h *Handler) Register(g *echo.Group) {
	v1 := g.Group("/v1")

	v1.POST("/shorten", h.ShortenUrl)
	v1.GET("/:shortUrlPath", h.GetOrgUrl)
}
