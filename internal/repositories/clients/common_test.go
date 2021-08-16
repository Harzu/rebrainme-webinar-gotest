package clients

import (
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"github.com/vrischmann/envconfig"

	"rebrainme/gotest/internal/config"
	"rebrainme/gotest/internal/system/database/psql"
)

const skipTestMessage = "Skip test. please up local database for this test"

func getTestRepoAndClient(req *require.Assertions) (*repositoryDB, *sqlx.DB) {
	type testConfig struct {
		PSQL *config.PSQL
	}

	var cfg testConfig
	err := envconfig.Init(&cfg)
	req.NoError(err)

	postgresClient, err := psql.New(cfg.PSQL)
	req.NoError(err)

	var (
		postgresConn = postgresClient.GetConnection()
		repository   = repositoryDB{conn: postgresConn}
	)

	return &repository, postgresConn
}
