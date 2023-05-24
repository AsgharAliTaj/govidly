package database

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func DatabseSetup() {
	var err error
	db, err = sqlx.Connect(
		"postgres",
		"user=root dbname=govidly password=toor port=5432 sslmode=disable",
	)

	if err != nil {
		log.Fatal("couldn't connect to db fool", err)
	}
	schema, err := os.ReadFile(".././database/Setup.sql")
	if err != nil {
		log.Fatal("schema failed to open", err)
	}
	db.MustExec(string(schema))
}

func InitDB() *sqlx.DB {
	return db
}
