package sqlmap

import (
	"database/sql"
)

type Result map[string]string
type ResultSet []Result

func Query(db *sql.DB, query string) (ResultSet, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	var results ResultSet

	for rows.Next() {
		result := make(Result)
		values := make([]string, len(cols))
		dest := make([]interface{}, len(cols))
		for i := range values {
			dest[i] = &values[i]
		}
		rows.Scan(dest...)
		for _, col := range cols {
			for _, v := range values {
				result[col] = v
			}
		}
		results = append(results, result)
	}

	return results, nil
}
