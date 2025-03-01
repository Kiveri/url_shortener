package sevice_provider

import (
	"net/http"

	urlshortener "url/internal/pb/url_shortener"

	"google.golang.org/grpc"
)

func (sp *ServiceProvider) GetHttpServer() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("POST /urls", sp.GetUrlsHttpController().CreateShortUrl)
	mux.HandleFunc("GET /", sp.GetUrlsHttpController().FindLongUrl)

	return mux
}

func (sp *ServiceProvider) GetGrpcServer() *grpc.Server {
	mux := grpc.NewServer()
	urlshortener.RegisterUrlShortenerServer(mux, sp.GetUrlsGrpcController())

	return mux
}
