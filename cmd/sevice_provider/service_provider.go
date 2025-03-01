package sevice_provider

import (
	"url/internal/adapter/storage/in_memory"
	"url/internal/adapter/storage/postgres"
	"url/internal/config"
	"url/internal/controller/grpc_controller"
	"url/internal/controller/http_controller"
	"url/internal/usecase"
)

type ServiceProvider struct {
	dbCluster         *config.Cluster
	urlUseCase        *usecase.UrlUseCase
	urlRepoInMem      *in_memory.Repo
	urlRepoDb         *postgres.Repo
	urlControllerHttp *http_controller.Controller
	urlControllerGrpc *grpc_controller.Controller
	conf              *config.Config
}

func NewServiceProvider(conf *config.Config) *ServiceProvider {
	return &ServiceProvider{
		conf: conf,
	}
}
