package handler

type getLongUrlRequest struct {
	ShortUrl string `param:"shortUrl" validate:"required"`
}

type shortenUrlRequest struct {
	LongUrl string `json:"longUrl" validate:"required"`
}
