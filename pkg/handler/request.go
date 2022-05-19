package handler

type getOrgUrlRequest struct {
	ShortUrlPath string `param:"shortUrlPath" validate:"required"`
}

type shortenUrlRequest struct {
	OrgUrl string `json:"orgUrl" validate:"required"`
}
