package sevice_provider

import (
	"url/internal/usecase"
)

func (sp *ServiceProvider) GetUrlUseCaseToDb() *usecase.UrlUseCase {
	if sp.urlUseCase == nil {
		sp.urlUseCase = usecase.NewUseCase(sp.getURLRepositoryDb())
	}

	return sp.urlUseCase
}

func (sp *ServiceProvider) GetUrlUseCaseToInMemory() *usecase.UrlUseCase {
	if sp.urlUseCase == nil {
		sp.urlUseCase = usecase.NewUseCase(sp.getURLRepositoryInMem())
	}

	return sp.urlUseCase
}
