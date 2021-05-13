package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello")
	})
	r.Run(":333")
}
