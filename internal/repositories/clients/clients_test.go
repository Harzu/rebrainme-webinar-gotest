package clients

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"rebrainme/gotest/internal/entities"
	"rebrainme/gotest/test/fixtures"
)

func TestRepositoryDB_InsertOrUpdateClient(t *testing.T) {
	if testing.Short() {
		t.Skip(skipTestMessage)
	}

	var (
		req              = require.New(t)
		ctx              = context.Background()
		repo, psqlClient = getTestRepoAndClient(req)
	)
	defer func() {
		req.NoError(psqlClient.Close())
	}()

	fixtures.ExecuteFixture(psqlClient, fixtures.CleanupFixture{})

	test := func(client entities.Client) func(t *testing.T) {
		return func(t *testing.T) {
			err := repo.InsertOrUpdateClient(ctx, client)
			req.NoError(err)

			var model clientModel
			row := psqlClient.QueryRowx(fmt.Sprintf("SELECT * FROM %s WHERE email = $1", tableClients), client.Email)
			err = row.StructScan(&model)
			req.NoError(err)
			req.Equal(client, buildClientEntity(model))
		}
	}

	clientEmail := "example@example.com"

	t.Run("insert new client", test(entities.Client{Email: clientEmail, FullName: "John Doe", City: "Moscow"}))
	t.Run("update exist client", test(entities.Client{Email: clientEmail, FullName: "John Braun", City: "NY"}))
}

func TestRepositoryDB_FindClientsByEmails(t *testing.T) {
	if testing.Short() {
		t.Skip(skipTestMessage)
	}

	var (
		req              = require.New(t)
		ctx              = context.Background()
		repo, psqlClient = getTestRepoAndClient(req)
	)
	defer func() {
		req.NoError(psqlClient.Close())
	}()

	fixtures.ExecuteFixture(psqlClient, fixtures.CleanupFixture{})

	clients := []entities.Client{
		{Email: "example_1@example.com", FullName: "John Doe", City: "Moscow"},
		{Email: "example_2@example.com", FullName: "Alan Walker", City: "NY"},
		{Email: "example_3@example.com", FullName: "Bob Marley", City: "London"},
	}
	for _, client := range clients {
		err := repo.InsertOrUpdateClient(ctx, client)
		req.NoError(err)
	}

	_, err := psqlClient.Exec(fmt.Sprintf("UPDATE %s SET deleted_at = NOW() WHERE email = $1", tableClients), clients[2].Email)
	req.NoError(err)

	test := func(emails []string, want []entities.Client) func(t *testing.T) {
		return func(t *testing.T) {
			actual, err := repo.FindClientsByEmails(ctx, emails)
			req.NoError(err)
			req.Len(actual, len(want))
			req.Equal(want, actual)
		}
	}

	t.Run("get exists clients", test([]string{clients[0].Email, clients[1].Email}, []entities.Client{clients[0], clients[1]}))
	t.Run("get deleted client", test([]string{clients[2].Email}, []entities.Client(nil)))
	t.Run("get all clients", test([]string{clients[0].Email, clients[1].Email, clients[2].Email}, []entities.Client{clients[0], clients[1]}))
}
