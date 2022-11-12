package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getHome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "Hello World",
	})
}

func GetHello(c *gin.Context) {
	html := template.HTML("")
	c.HTML(200, "ss", html)
}

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/", getHome)
	r.GET("/hello", GetHello)

	r.Run()
}
