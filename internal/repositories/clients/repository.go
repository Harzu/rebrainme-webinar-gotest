package clients

import (
	"context"

	"github.com/jmoiron/sqlx"

	"rebrainme/gotest/internal/entities"
	"rebrainme/gotest/internal/system/database/psql"
)

type Repository interface {
	InsertOrUpdateClient(ctx context.Context, client entities.Client) error
	FindClientByEmail(ctx context.Context, email string) (*entities.Client, error)
}

type repositoryDB struct {
	conn *sqlx.DB
}

func NewRepository(psqlClient psql.Client) Repository {
	return &repositoryDB{conn: psqlClient.GetConnection()}
}
