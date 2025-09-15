package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/", func(context *gin.Context) {
		context.String(200, "hello world")
	})
	route.Run(":8000")
}
