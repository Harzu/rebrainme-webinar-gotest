package clients

import (
	"database/sql"
	"time"

	"rebrainme/gotest/internal/entities"
)

type clientModel struct {
	Email     string       `db:"email"`
	FullName  string       `db:"full_name"`
	City      string       `db:"city"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func buildClientModel(entity entities.Client) clientModel {
	return clientModel{
		Email:     entity.Email,
		FullName:  entity.FullName,
		City:      entity.City,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: sql.NullTime{},
	}
}

func buildClientEntity(dbModel clientModel) entities.Client {
	return entities.Client{
		Email:    dbModel.Email,
		FullName: dbModel.FullName,
		City:     dbModel.City,
	}
}
