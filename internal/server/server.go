package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Server interface {
	RegisterRoutes()
	Start() error
}

type EchoServer struct {
	echo   *echo.Echo
	config ServerConfig
}

func (s *EchoServer) RegisterRoutes() {
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(fmt.Sprintf(":%d", s.config.Port)); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occured: %s", err)
		return err
	}
	return nil
}

func NewEchoServer(config ServerConfig) (Server, error) {
	server := echo.New()
	echoServer := &EchoServer{server, config}
	return echoServer, nil
}
