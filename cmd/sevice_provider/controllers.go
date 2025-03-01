package sevice_provider

import (
	"url/internal/controller/http_controller"
)

func (sp *ServiceProvider) GetUrlsController() *http_controller.Controller {
	if sp.urlController == nil {
		sp.urlController = http_controller.NewController(sp.GetUrlUseCase(), sp.conf.BaseUrl)
	}

	return sp.urlController
}
