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

type echoServer struct {
	config     Config
	echo       *echo.Echo
	routeGroup *echo.Group
}

func (s *echoServer) Start() error {
	if err := s.echo.Start(fmt.Sprintf(":%d", s.config.Port)); err != nil && err != http.ErrServerClosed {
		return err
	}
	return nil
}

func (s *echoServer) AddRoute(method, path string, handler echo.HandlerFunc) {
	if s.routeGroup != nil {
		s.routeGroup.Add(method, path, handler)
	} else {
		s.echo.Add(method, path, handler)
	}
}

func (s *echoServer) PrefixRoute(prefix string) {
	s.routeGroup = s.echo.Group(prefix)
}

func NewEchoServer(config Config) Server {
	return &echoServer{config, echo.New(), nil}
}
