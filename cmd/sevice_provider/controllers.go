package sevice_provider

import (
	"url/internal/controller/url_controller"
)

func (sp *ServiceProvider) GetUrlsControllerToDb() *url_controller.Controller {
	if sp.urlController == nil {
		sp.urlController = url_controller.NewController(sp.GetUrlUseCaseToDb())
	}

	return sp.urlController
}

func (sp *ServiceProvider) GetUrlControllerToInMem() *url_controller.Controller {
	if sp.urlController == nil {
		sp.urlController = url_controller.NewController(sp.GetUrlUseCaseToInMemory())
	}

	return sp.urlController
}
