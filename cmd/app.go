package main

import (
	"net/http"

	"url/cmd/sevice_provider"
	"url/internal/config"
)

func main() {
	sp := sevice_provider.NewServiceProvider(config.NewConfig())

	err := http.ListenAndServe(":8080", sp.GetRoutes())
	if err != nil {
		panic(err)
	}
}
