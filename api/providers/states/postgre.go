package states

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"log/slog"
	"rest-std-lib/mvp/providers"
	"rest-std-lib/mvp/providers/states"
	"strings"
)

type PostgresComponentStateProvider struct {
	url string
	db  *sql.DB
}

func (s *PostgresComponentStateProvider) Init(config providers.ProviderConfig) error {
	connStr := "user=postgres dbname=postgres sslmode=disable password= host=localhost"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		slog.Error("Cannot connect to do db", "str", connStr, "err", err)
	}
	s.db = db

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS Component (" +
		"name TEXT NOT NULL PRIMARY KEY," +
		"namespace TEXT NOT NULL, " +
		"constrains TEXT NOT NULL, " +
		"metadata TEXT NOT NULL," +
		"properties TEXT NOT NULL," +
		"body TEXT NOT NULL)")

	if err != nil {
		slog.Error("Cannot create table component", "err", err)
	}

	return nil
}

func (s *PostgresComponentStateProvider) Upsert(request states.UpsertRequest) error {
	var bodyAsMap = make(map[string]interface{})
	bodyAsJson, err := json.Marshal(request.Body)
	var entityType string
	if err != nil {
		slog.Error("Cannot marshal request body", "err", err)
		return err
	}
	err = json.Unmarshal(bodyAsJson, &bodyAsMap)
	if err != nil {
		slog.Error("Cannot unmarshal body into map", "err", err)
	}

	for k, v := range bodyAsMap {
		if strings.EqualFold("type", k) || strings.EqualFold("kind", k) {
			entityType = v.(string)
		}
	}

	var queryStr = "INSERT INTO " + entityType + " VALUES ("

	for _, v := range bodyAsMap {
		if vv, ok := v.(string); ok {
			queryStr = queryStr + "'" + vv + "',"
		} else {
			vv, err := json.Marshal(v)
			if err != nil {
				queryStr = queryStr + "'{}',"
				slog.Warn("Failure in marshaling", "obj", vv)
			}
			queryStr = queryStr + "'" + string(vv) + "',"
		}
	}

	queryStr = queryStr + "'" + string(bodyAsJson) + "'" + ")"
	_, err = s.db.Exec(queryStr)

	return nil
}

func (s *PostgresComponentStateProvider) List(request states.ListRequest) []states.StateEntry {
	rows, err := s.db.Query("SELECT * FROM Component")
	if err != nil {
		slog.Error("Error while select all from component")
		return nil
	}

	cols, _ := rows.Columns()
	var results []states.StateEntry
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return nil
		}

		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		result := states.StateEntry{
			ID:   m["name"].(string),
			Body: []byte(m["body"].(string)),
		}
		results = append(results, result)
	}

	return results
}
func (s *PostgresComponentStateProvider) Get(request states.GetRequest) (states.StateEntry, error) {
	rows, err := s.db.Query("SELECT * FROM Component WHERE name=" + "'" + request.ID + "'")

	if err != nil {
		slog.Error("Error while select by name")
		return states.StateEntry{}, nil
	}

	cols, _ := rows.Columns()
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			return states.StateEntry{}, nil
		}

		m := make(map[string]interface{})
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
		result := states.StateEntry{
			ID:   m["name"].(string),
			Body: []byte(m["body"].(string)),
		}
		return result, nil
	}

	return states.StateEntry{}, nil
}
