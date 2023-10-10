package postgres

import (
	"errors"
	"flag"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func New() (*sqlx.DB, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbHost := flag.String("dbhost", os.Getenv("DB_HOST"), "Database host")
	dbName := flag.String("dbname", os.Getenv("DB_NAME"), "Database name")
	dbUser := flag.String("dbuser", os.Getenv("DB_USER"), "Database user")
	dbPass := flag.String("dbpass", os.Getenv("DB_PASSWORD"), "Database password")
	dbPort := flag.String("dbport", os.Getenv("DB_PORT"), "Database port")
	dbSSL := flag.String("dbssl", os.Getenv("DB_SSL"), "Database ssl settings (disable, prefer, require)")

	connectionString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=%s", *dbHost, *dbPort, *dbName, *dbUser, *dbPass, *dbSSL)

	conn, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	err = conn.Ping()
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return conn, nil
}
