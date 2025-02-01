package sevice_provider

import (
	"net/http"
)

func (sp *ServiceProvider) GetRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /urls", sp.GetUrlsController().CreateShortUrl)
	mux.HandleFunc("GET /urls/{short_url}", sp.GetUrlsController().FindLongUrl)

	return mux
}
