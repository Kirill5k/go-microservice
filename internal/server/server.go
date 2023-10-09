package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Server interface {
	Start() error
	PrefixRoute(prefix string)
	AddRoute(method, path string, handler echo.HandlerFunc)
}

type RouteRegister interface {
	RegisterRoutes(server Server)
}

type EchoServer struct {
	config     Config
	echo       *echo.Echo
	routeGroup *echo.Group
}

func (s *EchoServer) Start() error {
	if err := s.echo.Start(fmt.Sprintf(":%d", s.config.Port)); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *EchoServer) AddRoute(method, path string, handler echo.HandlerFunc) {
	if s.routeGroup != nil {
		s.routeGroup.Add(method, path, handler)
	} else {
		s.echo.Add(method, path, handler)
	}
}

func (s *EchoServer) PrefixRoute(prefix string) {
	s.routeGroup = s.echo.Group(prefix)
}

func NewEchoServer(config Config) *EchoServer {
	return &EchoServer{config, echo.New(), nil}
}
