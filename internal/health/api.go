package health

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
	"kirill5k/go/microservice/internal/database"
	"kirill5k/go/microservice/internal/server"
	"net"
	"net/http"
	"os"
	"time"
)

type Api struct {
	dbClient    database.Client
	ipAddress   string
	appVersion  string
	startupTime time.Time
}

func (api *Api) RegisterRoutes(server server.Server) {
	server.PrefixRoute("/health")

	server.AddRoute("GET", "/ready", func(ctx echo.Context) error {
		if api.dbClient.Ready() {
			return ctx.JSON(http.StatusOK, StatusUp(api.startupTime, api.ipAddress, api.appVersion))
		}
		return ctx.JSON(http.StatusServiceUnavailable, StatusDown(api.startupTime, api.ipAddress, api.appVersion))
	})

	server.AddRoute("GET", "/live", func(ctx echo.Context) error {
		return ctx.JSON(http.StatusOK, StatusUp(api.startupTime, api.ipAddress, api.appVersion))
	})
}

func NewApi(dbClient database.Client) *Api {
	getIpaddress := func() string {
		conn, err := net.Dial("udp", "8.8.8.8:80")
		if err != nil {
			log.Fatal().Err(err).Msg("failed to obtain ip address")
		}
		defer func(conn net.Conn) {
			err := conn.Close()
			if err != nil {
				log.Fatal().Err(err).Msg("failed to obtain ip address")
			}
		}(conn)

		localAddr := conn.LocalAddr().(*net.UDPAddr)

		return localAddr.String()
	}

	return &Api{
		dbClient:    dbClient,
		ipAddress:   getIpaddress(),
		appVersion:  os.Getenv("VERSION"),
		startupTime: time.Now(),
	}
}
