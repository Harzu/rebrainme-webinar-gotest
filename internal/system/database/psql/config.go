package psql

type Config struct {
	DSN         string
	MaxOpenConn int `envconfig:"default=0"`
}
