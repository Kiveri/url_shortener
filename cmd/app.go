package main

import (
	"flag"
	"fmt"
	"net/http"

	"url/cmd/sevice_provider"
	"url/internal/config"
)

func main() {
	storageType := flag.String("storage", "postgresql", "Тип хранилища (in-memory или postgresql)")
	flag.Parse()
	var c *config.Config

	switch *storageType {
	case "in-memory":
		c = config.NewConfig(*storageType)
	case "postgresql":
		c = config.NewConfig(*storageType)
	default:
		panic("данный тип не але")
	}
	fmt.Println(c.DatabaseType)

	sp := sevice_provider.NewServiceProvider(c)

	err := http.ListenAndServe(":8080", sp.GetRoutes())
	if err != nil {
		panic(err)
	}
}
