package sevice_provider

import (
	"context"

	"url/internal/adapter/storages/in_memory"
	"url/internal/adapter/storages/postgres"
)

func (sp *ServiceProvider) getURLRepositoryInMem() *in_memory.Repo {
	if sp.urlRepoInMem == nil {
		sp.urlRepoInMem = in_memory.NewRepo()
	}

	return sp.urlRepoInMem
}

func (sp *ServiceProvider) getURLRepositoryDb() *postgres.Repo {
	if sp.urlRepoDb == nil {
		sp.urlRepoDb = postgres.NewRepoDb(sp.getDbCluster(context.Background()))
	}

	return sp.urlRepoDb
}
