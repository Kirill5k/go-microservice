package customer

import "kirill5k/go/microservice/internal/database"

type Module struct {
	Api *Api
}

func NewModule(client *database.PostgresClient) *Module {
	repository := NewPostgresRepository(client)
	service := NewService(repository)
	api := NewApi(service)
	return &Module{api}
}
