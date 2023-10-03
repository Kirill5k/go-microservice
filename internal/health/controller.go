package health

import (
	"github.com/labstack/echo/v4"
	"kirill5k/go/microservice/internal/database"
	"kirill5k/go/microservice/internal/server"
	"net/http"
)

type HealthController struct {
	dbClient database.Client
}

func (hc *HealthController) RegisterRoutes(server server.Server) {
	readiness := func(ctx echo.Context) error {
		isReady := hc.dbClient.Ready()
		if isReady {
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

func NewHealthController(dbClient database.Client) *HealthController {
	return &HealthController{dbClient}
}
