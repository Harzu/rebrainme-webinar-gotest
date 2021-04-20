package fixtures

import "github.com/jmoiron/sqlx"

type Fixture interface {
	GetSql() []string
}

type CleanupFixture struct{}

func (cf CleanupFixture) GetSql() []string {
	return []string{
		`TRUNCATE TABLE clients RESTART IDENTITY CASCADE;`,
	}
}

func ExecuteFixture(client *sqlx.DB, fixture Fixture) {
	for _, query := range fixture.GetSql() {
		_, err := client.Exec(query)

		if err != nil {
			panic(err)
		}
	}
}
