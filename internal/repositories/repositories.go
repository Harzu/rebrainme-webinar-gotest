package repositories

import (
	"rebrainme/gotest/internal/repositories/clients"
	"rebrainme/gotest/internal/system/database/psql"
)

type Container struct {
	Clients clients.Repository
}

func New(psqlClient psql.Client) *Container {
	return &Container{
		Clients: clients.NewRepository(psqlClient),
	}
}
