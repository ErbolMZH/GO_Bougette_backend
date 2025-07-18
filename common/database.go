package common

import (
	"fmt"
	"github.com/joho/godotenv"
	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func NewMysql() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, username, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Default().Println("Database connection established")
	return db, nil

	//fmt.Println(dsn)
}
