package health

import (
	"github.com/labstack/echo/v4"
	"kirill5k/go/microservice/internal/database"
	"kirill5k/go/microservice/internal/server"
	"net/http"
)

type Api struct {
	dbClient database.Client
}

func (hc *Api) RegisterRoutes(server server.Server) {
	readiness := func(ctx echo.Context) error {
		if hc.dbClient.Ready() {
			return ctx.JSON(http.StatusOK, Status{"OK"})
		}
		return ctx.JSON(http.StatusServiceUnavailable, Status{"NOT_AVAILABLE"})
	}
	server.AddRoute("GET", "/health/ready", readiness)

	liveness := func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, Status{"OK"})
	}
	server.AddRoute("GET", "/health/live", liveness)
}

func NewApi(dbClient database.Client) *Api {
	return &Api{dbClient}
}
