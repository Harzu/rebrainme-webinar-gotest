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

func (r *repositoryDB) FindClientsByEmails(ctx context.Context, emails []string) (result []entities.Client, err error) {
	query, args, err := prepareFindClientsByEmailsQuery(emails)
	if err != nil {
		return result, fmt.Errorf("failed to prepare FindClientsByEmails query: %w", err)
	}

	rows, err := r.conn.QueryxContext(ctx, query, args...)
	if err != nil {
		return result, fmt.Errorf("failed to execute FindClientsByEmails query: %w", err)
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	for rows.Next() {
		var model clientModel
		if err := rows.StructScan(&model); err != nil {
			return result, fmt.Errorf("failed to scan row to struct")
		}
		result = append(result, buildClientEntity(model))
	}

	if err := rows.Err(); err != nil {
		return result, fmt.Errorf("unable to scan all out of FindClientsByEmails query: %w", err)
	}

	return
}
