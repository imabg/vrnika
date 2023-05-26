package database

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// create connection to postgres db
func ConnectToDB() error {
	portStr := os.Getenv("DB_PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PWD"), os.Getenv("DB_NAME"), port)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
