package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", getPing)
	r.POST("/ping", postPing)

	r.GET("/detail/:id", getDetail)

	r.Run(":333")
}

func getDetail(context *gin.Context) {
	id := context.Param("id")
	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func getPing(context *gin.Context) {
	name := context.DefaultQuery("name", "Guest")
	var data = map[string]interface{}{
		"messeage": "Hello " + name + " from Getping",
	}
	context.JSON(http.StatusOK, data)

}
func postPing(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"Tung": "Nui",
	})
}
