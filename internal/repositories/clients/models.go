package clients

import (
	"database/sql"
	"time"
)

type clientModel struct {
	Email     string       `db:"email"`
	FullName  string       `db:"full_name"`
	City      string       `db:"city"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
