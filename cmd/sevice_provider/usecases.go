package sevice_provider

import (
	"url/internal/usecase"
)

func (sp *ServiceProvider) GetUrlUseCase() *usecase.UrlUseCase {
	if sp.urlUseCase == nil {
		if sp.conf.DatabaseType == "in-memory" {
			sp.urlUseCase = usecase.NewUseCase(sp.getURLRepositoryInMem())
		}
		if sp.conf.DatabaseType == "postgresql" {
			sp.urlUseCase = usecase.NewUseCase(sp.getURLRepositoryDb())
		}
	}

	return sp.urlUseCase
}
