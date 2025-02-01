package main

import (
	"net/http"

	"url/cmd/sevice_provider"
)

func main() {
	sp := sevice_provider.NewServiceProvider()

	err := http.ListenAndServe(":8080", sp.GetRoutes())
	if err != nil {
		panic(err)
	}
}
