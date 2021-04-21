package clients

import (
	sq "github.com/Masterminds/squirrel"

	"rebrainme/gotest/internal/system/database/stommer"
)

const tableClients = "clients"

func prepareInsertOrUpdateClientQuery(model clientModel) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	st, err := stommer.New(model)
	if err != nil {
		return "", []interface{}{}, err
	}

	rawRequest := psqlSq.Insert(tableClients).
		Columns(st.Columns...).
		Values(st.Values...).
		Suffix(`
			ON CONFLICT (email) DO UPDATE SET
				full_name  = EXCLUDED.full_name,
				city       = EXCLUDED.city,
				updated_at = NOW(),
				deleted_at = NULL
		`)

	return rawRequest.ToSql()
}

func prepareFindClientByEmailQuery(email string) (string, []interface{}, error) {
	psqlSq := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	st, err := stommer.New(clientModel{})
	if err != nil {
		return "", []interface{}{}, err
	}

	rawRequest := psqlSq.Select(st.Columns...).
		From(tableClients).
		Where(sq.Eq{
			"email":      email,
			"deleted_at": nil,
		})

	return rawRequest.ToSql()
}
