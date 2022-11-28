package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title  string
	Author string
	Price  string
}

func AllBoks(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Atomic Habits",
	})
}

func StoreBook(c *gin.Context) {
	var book Book
	err := c.ShouldBindJSON(&book)
	if err != nil {
		log.Fatal("Failed bind json")
	}

	c.JSON(200, gin.H{
		"Title": book.Title,
		"Price": book.Price,
	})

}
