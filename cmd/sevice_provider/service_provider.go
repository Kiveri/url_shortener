package sevice_provider

import (
	"url/internal/adapter/storages/in_memory"
	"url/internal/adapter/storages/postgres"
	"url/internal/config"
	"url/internal/controller/url_controller"
	"url/internal/usecase"
)

type ServiceProvider struct {
	dbCluster *config.Cluster

	urlUseCase *usecase.UrlUseCase

	urlRepoInMem *in_memory.Repo
	urlRepoDb    *postgres.Repo

	urlController *url_controller.Controller
}

func NewServiceProvider() *ServiceProvider {
	return &ServiceProvider{}
}
