package customer

import "kirill5k/go/microservice/internal/server"

type Api struct {
	service *Service
}

func (hc *Api) RegisterRoutes(server server.Server) {

}

func NewApi(service *Service) *Api {
	return &Api{service}
}
