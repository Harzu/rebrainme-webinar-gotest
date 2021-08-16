package config

type PSQL struct {
	DSN         string `envconfig:"PSQL_DSN"`
	MaxOpenConn int    `envconfig:"PSQL_MAX_OPEN_CONN,default=0"`
}
