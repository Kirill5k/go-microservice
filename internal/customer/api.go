package customer

import (
	"github.com/labstack/echo/v4"
	"kirill5k/go/microservice/internal/server"
	"net/http"
)

type Api struct {
	service *Service
}

func (hc *Api) RegisterRoutes(server server.Server) {
	server.PrefixRoute("/customers")
	getAll := func(ctx echo.Context) error {
		email := ctx.QueryParam("email")
		customers, err := hc.service.FindBy(ctx.Request().Context(), email)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}
		return ctx.JSON(http.StatusOK, customers)
	}
	server.AddRoute("GET", "", getAll)
}

func NewApi(service *Service) *Api {
	return &Api{service}
}
