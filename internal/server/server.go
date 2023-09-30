package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"kirill5k/go/microservice/internal/database"
	"log"
	"net/http"
)

type Server interface {
	Start(config ServerConfig) error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

func (s *EchoServer) Start(config ServerConfig) error {
	if err := s.echo.Start(fmt.Sprintf(":%d", config.Port)); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occured: %s", err)
		return err
	}
	return nil
}

func (s *EchoServer) registerRoutes() {

}

func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{echo.New(), db}
	server.registerRoutes()
	return server
}
