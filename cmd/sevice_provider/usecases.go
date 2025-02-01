package sevice_provider

import (
	"url/internal/usecase"
)

func (sp *ServiceProvider) GetUrlUseCase() *usecase.UrlUseCase {
	if sp.urlUseCase == nil {
		sp.urlUseCase = usecase.NewUseCase(sp.getURLRepositoryDb())
	}

	return sp.urlUseCase
}
