package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Server interface {
	RegisterRoutes()
}

type EchoServer struct {
	echo *echo.Echo
}

func (s *EchoServer) RegisterRoutes() {
}

func NewEchoServer(config ServerConfig) (Server, error) {
	server := echo.New()
	if err := server.Start(fmt.Sprintf(":%d", config.Port)); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occured: %s", err)
		return nil, err
	}
	echoServer := &EchoServer{server}
	return echoServer, nil
}
