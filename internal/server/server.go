package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type Server interface {
	Start() error
	AddRoute(method, path string, handler echo.HandlerFunc)
}

type EchoServer struct {
	echo   *echo.Echo
	config ServerConfig
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(fmt.Sprintf(":%d", s.config.Port)); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occured: %s", err)
		return err
	}
	return nil
}

func (s *EchoServer) AddRoute(method, path string, handler echo.HandlerFunc) {
	s.echo.Add(method, path, handler)
}

func NewEchoServer(config ServerConfig) (Server, error) {
	server := &EchoServer{echo.New(), config}
	return server, nil
}
