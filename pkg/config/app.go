package config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func Connect() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env files")
		return
	}
	host := os.Getenv("DB_DEV_HOST")
	port := os.Getenv("DB_DEV_PORT")
	dbname := os.Getenv("DB_DEV_NAME")
	user := os.Getenv("DB_DEV_USER")
	password := os.Getenv("DB_DEV_PASSWORD")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	d, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
