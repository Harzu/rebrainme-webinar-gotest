package clients

import (
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/vrischmann/envconfig"

	"rebrainme/gotest/internal/system/database/psql"
)

const skipTestMessage = "Skip test. please up local database for this test"

func getTestRepoAndClient(req *require.Assertions) (Repository, *sqlx.DB) {
	type config struct {
		PSQL *psql.Config
	}

	var cfg config
	err := envconfig.Init(&cfg)
	req.NoError(err)

	postgresClient, err := psql.New(cfg.PSQL)
	req.NoError(err)

	repo := NewRepository(postgresClient)

	return repo, postgresClient.GetConnection()
}
