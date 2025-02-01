package sevice_provider

import (
	"net/http"
)

func (sp *ServiceProvider) GetRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /urls", sp.GetUrlsControllerToDb().CreateShortUrl)
	mux.HandleFunc("GET /urls/{short_url}", sp.GetUrlsControllerToDb().FindLongUrl)

	//mux.HandleFunc("POST /urls", sp.GetUrlControllerToInMem().CreateShortUrl)
	//mux.HandleFunc("GET /urls/{short_url}", sp.GetUrlControllerToInMem().FindLongUrl)

	return mux
}
