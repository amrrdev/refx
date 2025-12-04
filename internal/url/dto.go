package url

type CreateShortUrlBody struct {
	LongUrl string `json:"long_url"`
}

type RedirectBody struct {
	ShortUrl string `json:"short_url"`
}
