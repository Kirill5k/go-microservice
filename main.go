package main

import (
	"kirill5k/go/microservice/internal/database"
	"kirill5k/go/microservice/internal/health"
	"kirill5k/go/microservice/internal/server"
	"log"
)

func main() {
	db, err := database.NewPostgresClient(database.DefaultPostgresConfig())
	if err != nil {
		log.Fatalf("failed to initialise postgres client: %s", err)
	}
	srv := server.NewEchoServer(server.DefaultConfig())

	healthController := health.NewHealthController(db)
	healthController.RegisterRoutes(srv)

	if err := srv.Start(); err != nil {
		log.Fatalf("failed to start http server: %s", err)
	}
}
