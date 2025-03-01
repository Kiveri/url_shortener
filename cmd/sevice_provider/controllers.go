package sevice_provider

import (
	"url/internal/controller/grpc_controller"
	"url/internal/controller/http_controller"
)

func (sp *ServiceProvider) GetUrlsHttpController() *http_controller.Controller {
	if sp.urlControllerHttp == nil {
		sp.urlControllerHttp = http_controller.NewController(sp.GetUrlUseCase(), sp.conf.BaseUrl)
	}

	return sp.urlControllerHttp
}

func (sp *ServiceProvider) GetUrlsGrpcController() *grpc_controller.Controller {
	if sp.urlControllerGrpc == nil {
		sp.urlControllerGrpc = grpc_controller.NewController(sp.GetUrlUseCase(), sp.conf.BaseUrl)
	}

	return sp.urlControllerGrpc
}
