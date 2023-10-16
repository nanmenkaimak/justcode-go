package postgres

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/nanmenkaimak/justcode-go/lecture_9/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func New() (*gorm.DB, error) {
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

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connectionString,
		PreferSimpleProtocol: true,
	}))

	db.AutoMigrate(&models.Gist{}, &models.Commit{}, &models.File{})

	return db, nil
}
