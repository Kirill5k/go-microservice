package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server interface {
	Start() error
	AddRoute(method, path string, handler echo.HandlerFunc)
}

type EchoServer struct {
	echo   *echo.Echo
	config Config
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(fmt.Sprintf(":%d", s.config.Port)); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *EchoServer) AddRoute(method, path string, handler echo.HandlerFunc) {
	s.echo.Add(method, path, handler)
}

func NewEchoServer(config Config) Server {
	server := &EchoServer{echo.New(), config}
	return server
}
