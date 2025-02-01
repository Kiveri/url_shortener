package usecase

type urlRepo interface {
	CreateShortUrl(shortUrl, longURL string) (string, error)
	FindByShortUrl(shortUrl string) (string, error)
}

type UrlUseCase struct {
	urlRepo urlRepo
}

func NewUseCase(urlRepo urlRepo) *UrlUseCase {
	return &UrlUseCase{
		urlRepo: urlRepo,
	}
}
