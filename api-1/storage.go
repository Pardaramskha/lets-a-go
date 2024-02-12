package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connectionCreds := "user=postgres dbname=postgres password=password123 sslmode=disable"
	db, err := sql.Open("postgres", connectionCreds)
	// check connection launch
	if err != nil {
		return nil, err
	}
	// ping DB and check for errors
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (store *PostgresStore) InitTodoTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS TODOS1 (
			id serial primary key,
			label TEXT,
			checked BOOLEAN,
			created_at TEXT
		)`

	_, err := store.db.Exec(query)
	if err != nil {
		panic(err)
	}

	return nil

	//query2 := `
	//	INSERT INTO todos1 (ID, LABEL, CHECKED, CREATED_AT)
	//	VALUES ('01', 'test', false, 'now')`
	//
	//_, err2 := store.db.Exec(query2)
	//return err2
}
