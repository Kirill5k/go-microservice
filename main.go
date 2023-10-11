package main

import (
	"kirill5k/go/microservice/internal/config"
	"kirill5k/go/microservice/internal/customer"
	"kirill5k/go/microservice/internal/database"
	"kirill5k/go/microservice/internal/health"
	"kirill5k/go/microservice/internal/server"
	"log"
)

func main() {
	conf := config.LoadViperConfig()
	db, err := database.NewPostgresClient(&conf.Postgres)
	if err != nil {
		log.Fatalf("failed to initialise postgres client: %s", err)
	}
	srv := server.NewEchoServer(&conf.Server)

	apis := []server.RouteRegister{
		health.NewModule(db).Api,
		customer.NewModule(db).Api,
	}
	for _, api := range apis {
		api.RegisterRoutes(srv)
	}

	if err := srv.Start(); err != nil {
		log.Fatalf("failed to start http server: %s", err)
	}
}
