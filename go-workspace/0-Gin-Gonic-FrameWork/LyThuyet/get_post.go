package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", getPing)
	r.POST("/ping", postPing)

	r.Run(":333")
}

func getPing(context *gin.Context) {
	context.String(http.StatusOK, "Hello")
}
func postPing(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Tung": "Nui",
	})
}
