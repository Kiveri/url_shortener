package main

import (
	"fmt"
	"net"
	"net/http"

	"url/cmd/sevice_provider"
	"url/internal/config"
)

func main() {
	sp := sevice_provider.NewServiceProvider(config.NewConfig())

	go func() {
		err := http.ListenAndServe(":8080", sp.GetHttpServer())
		if err != nil {
			panic(err)
		}
	}()

	grpcServer := sp.GetGrpcServer()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if con := grpcServer.Serve(lis); con != nil {
		fmt.Printf("failed to serve: %v\n", con)
	}
}
