package clients

import (
	"database/sql"
	"time"

	"rebrainme/gotest/internal/entities"
)

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
