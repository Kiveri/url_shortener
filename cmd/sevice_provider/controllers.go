package sevice_provider

import (
	"url/internal/controller/url_controller"
)

func (sp *ServiceProvider) GetUrlsController() *url_controller.Controller {
	if sp.urlController == nil {
		sp.urlController = url_controller.NewController(sp.GetUrlUseCase())
	}

	return sp.urlController
}
