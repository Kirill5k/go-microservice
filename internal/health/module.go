package health

import "kirill5k/go/microservice/internal/database"

type Module struct {
	Api *Api
}

func NewModule(dbClient database.Client) *Module {
	api := NewApi(dbClient)
	return &Module{api}
}
