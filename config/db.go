package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect with the db")
	}
	DB = db
	log.Println("Connected to the Db")

}
