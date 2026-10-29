package pgsql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

func Init() (*sqlx.DB, error) {
	dbUser := os.Getenv("DATABASE_USER")
	dbName := os.Getenv("DATABASE_NAME")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")

	if dbUser == "" || dbName == "" || dbPass == "" || dbHost == "" || dbPort == "" {
		return nil, fmt.Errorf("database configuration is not set properly")
	}

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPass, dbName, dbHost, dbPort)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("Successfully connected to the database")
	return db, nil
}
