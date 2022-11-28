package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hadigun007/golang-api/handler"
	initial "github.com/hadigun007/golang-api/init"
)

func init() {
	initial.LoadEnv()
	initial.LoadDB()
}

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/books", handler.AllBoks)
	r.POST("/books", handler.StoreBook)

	r.Run()
}
