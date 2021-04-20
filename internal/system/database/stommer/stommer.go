package stommer

import (
	"database/sql/driver"
	"fmt"

	"github.com/elgris/stom"
)

type stommer struct {
	mapper  map[string]interface{}
	Columns []string
	Values  []interface{}
}

func New(o interface{}, omitted ...string) (*stommer, error) {
	m, err := stom.MustNewStom(o).ToMap(o)
	if err != nil {
		return nil, err
	}

	var (
		columns = make([]string, 0, len(m)-len(omitted))
		values  = make([]interface{}, 0, len(m)-len(omitted))
	)

	for k, v := range m {
		if checkOmitted(k, omitted...) {
			continue
		}
		columns = append(columns, k)
		tmpVal := v
		if valuer, ok := v.(driver.Valuer); ok {
			tmpVal, err = valuer.Value()
			if err != nil {
				return nil, fmt.Errorf("could not convert value: %w", err)
			}
		}

		values = append(values, tmpVal)
	}

	return &stommer{
		mapper:  m,
		Columns: columns,
		Values:  values,
	}, nil
}

func checkOmitted(key string, omitted ...string) bool {
	for _, omitted := range omitted {
		if key == omitted {
			return true
		}
	}

	return false
}
