package http_controller

import (
	"errors"
	"net/http"

	"url/internal/adapter/storage"
	"url/internal/controller"
	"url/internal/usecase"
)

type findLongUrlResponse struct {
	LongUrl string `json:"long_url"`
}

func (c *HttpController) FindLongUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Path[1:]
	longUrl, err := c.urlsUseCase.FindUrl(usecase.FindUrlReq{
		ShortUrl: shortUrl,
	})
	if err != nil {
		if errors.Is(err, storage.NotFound) {
			NotFoundErrorRespond(w, controller.NewNotFoundError("not found"))

			return
		}

		InternalServerErrorRespond(w, err)

		return
	}

	Respond(w, http.StatusOK, findLongUrlResponse{LongUrl: longUrl.LongUrl})
}
