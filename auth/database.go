package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// CreateConnection - ...
func CreateConnection() (*sql.DB, error) {
	host := "localhost"                   // os.Getenv("DB_HOST")
	databaseUser := "postgres"            // os.Getenv("DB_USER")
	databaseName := "simple_auth"         // os.Getenv("DB_NAME")
	databasePassword := "postgres-simple" // os.Getenv("DB_PASSWORD")

	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=disable password=%s",
			host, databaseUser, databaseName, databasePassword,
		),
	)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	return db, err
}
