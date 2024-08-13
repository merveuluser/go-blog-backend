package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var DB *sql.DB

func InitDB() {
	var err error

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	if username == "" || password == "" || dbname == "" || sslmode == "" {
		log.Fatal("DB_USER, DB_PASSWORD, DB_NAME and DB_SSLMODE are required. Check your env variables.")
	}

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s", username, password, dbname, sslmode)

	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("sql.Open failed: ", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("DB.Ping failed: ", err)
	}

	log.Println("Successfully connected to database")
}

func CreateTables() error {
	databaseFilePath := os.Getenv("CREATE_TABLES_FILEPATH")

	data, err := os.ReadFile(databaseFilePath)
	if err != nil {
		log.Fatal("Unable to read .sql file: ", err)
	}

	queries := string(data)

	_, err = DB.Exec(queries)
	return err
}

func DeleteTables() error {
	databaseFilePath := os.Getenv("DELETE_TABLES_FILEPATH")

	data, err := os.ReadFile(databaseFilePath)
	if err != nil {
		log.Fatal("Unable to read .sql file", err)
	}

	queries := string(data)

	_, err = DB.Exec(queries)
	return err
}
