package model

type URL struct {
	ShortUrl string
	LongUrl  string
}

func NewURL(shortUrl, longURL string) *URL {
	return &URL{
		ShortUrl: shortUrl,
		LongUrl:  longURL,
	}
}
