package storage

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func setDatabaseConnectionString() string {
	err := godotenv.Load()
	if err != nil {
		panic("No environment found!")
	}

	return fmt.Sprintf("user=%v dbname=%v password=%v sslmode=disable", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_PWD"))
}

func NewPostgresStore() *sql.DB {
	connectionCreds := setDatabaseConnectionString()
	db, err := sql.Open("postgres", connectionCreds)
	// check connection launch
	if err != nil {
		log.Fatalf("Fatal error on database opening: %v", err)
	}
	// ping DB and check for errors
	if err := db.Ping(); err != nil {
		log.Fatalf("Fatal error on database ping: %v", err)
	}

	return db
}
