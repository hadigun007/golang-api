package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	MigrateBook()
	MigrateUser()
}

func MigrateUser() {

	type UserModel struct {
		gorm.Model
		Title string
		Price string
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed Load .env")
	}
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	DB.AutoMigrate(&UserModel{})
}

func MigrateBook() {

	type BookModel struct {
		gorm.Model
		Title string
		Price string
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed Load .env")
	}
	dsn := os.Getenv("DB_URL")
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	DB.AutoMigrate(&BookModel{})
}
