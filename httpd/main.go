package main

import (
	"golandemo/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.GET("/greet", handler.Greet)
	r.Run(":8000")
}
