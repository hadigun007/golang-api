package initial

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env")
	}
}

var DB gorm.DB

func LoadDB() {
	var errs error
	dsn := os.Getenv("DB_URL")
	DB, errs := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if errs != nil {
		log.Fatal("DB  Connection error", errs)
		fmt.Println(DB)
	}
}
