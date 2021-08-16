package psql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"rebrainme/gotest/internal/config"
)

const dialect = "postgres"

type Client interface {
	GetConnection() *sqlx.DB
}

type client struct {
	conn *sqlx.DB
}

func New(cfg *config.PSQL) (Client, error) {
	db, err := sqlx.Open(dialect, cfg.DSN)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.MaxOpenConn)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &client{conn: db}, nil
}

func (c client) GetConnection() *sqlx.DB {
	return c.conn
}
