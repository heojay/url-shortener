package handler

type getLongUrlResponse struct {
	LongUrl string `json:"longUrl"`
}

type shortenUrlResponse struct {
	ShortUrl string `json:"shortUrl"`
}
