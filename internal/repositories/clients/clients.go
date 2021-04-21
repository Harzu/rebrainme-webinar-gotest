package clients

import (
	"context"
	"fmt"

	"rebrainme/gotest/internal/entities"
)

func (r *repositoryDB) InsertOrUpdateClient(ctx context.Context, client entities.Client) error {
	query, args, err := prepareInsertOrUpdateClientQuery(buildClientModel(client))
	if err != nil {
		return fmt.Errorf("failed to prepare InsertOrUpdateClient query: %w", err)
	}

	if _, err := r.conn.ExecContext(ctx, query, args...); err != nil {
		return fmt.Errorf("failed to execute InsertOrUpdateClient query: %w", err)
	}

	return nil
}

func (r *repositoryDB) FindClientByEmail(ctx context.Context, email string) (*entities.Client, error) {
	query, args, err := prepareFindClientByEmailQuery(email)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare FindClientByEmail query: %w", err)
	}

	var model clientModel
	if err := r.conn.QueryRowxContext(ctx, query, args...).StructScan(&model); err != nil {
		return nil, fmt.Errorf("failed to execute FindClientByEmail query: %w", err)
	}

	result := buildClientEntity(model)
	return &result, nil
}
