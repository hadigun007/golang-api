package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title     string
	Price     string
	Referensi []string
}

var books []Book
var book = Book{}
var ref []string

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed Load .env")
	}
	// dsn := os.Getenv("DB_URL")
	// DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// DB.AutoMigrate(&model.BookModel{})
	r := gin.Default()
	r.GET("/books", bookHandler)
	r.GET("/book/:id", findBook)
	r.GET("/book", queryBook)

	r.Run(":0909")
}

var fef []string

func findBook(ctx *gin.Context) {
	id := ctx.Param("id")
	fef = append(fef, id)

	ctx.JSON(200, gin.H{
		"data": fef,
	})
}

func queryBook(ctx *gin.Context) {

	name := ctx.Query("name")

	fef = append(fef, name)
	ctx.JSON(200, gin.H{
		"data": fef,
	})
}

func bookHandler(ctx *gin.Context) {
	ref = append(ref, "Baik", "Mantap")

	book.ID = 1
	book.Price = "300000"
	book.Title = "Bukan Doa"
	book.Referensi = ref

	books = append(books, book)
	ctx.JSON(200, gin.H{
		"msg": books,
	})
}
