package sevice_provider

import (
	"url/internal/adapter/storage/in_memory"
	"url/internal/adapter/storage/postgres"
	"url/internal/config"
	"url/internal/controller/http_controller"
	"url/internal/usecase"
)

type ServiceProvider struct {
	dbCluster     *config.Cluster
	urlUseCase    *usecase.UrlUseCase
	urlRepoInMem  *in_memory.Repo
	urlRepoDb     *postgres.Repo
	urlController *http_controller.HttpController
	conf          *config.Config
}

func NewServiceProvider(conf *config.Config) *ServiceProvider {
	return &ServiceProvider{
		conf: conf,
	}
}
