package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(ACustomMiddleware)

	r.GET("/ping", getPing)
	r.POST("/ping", postPing)

	r.GET("/detail/:id", getDetail)

	api := r.Group("/api")
	{
		v1 := api.Group("v1")
		{
			// Khởi tạo middleware chi dùng cho gr v1
			v1.Use(MyGroupV1Middleware())
			v1.GET("/ping", func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{
					"messeage": "ping",
				})
			})

			v1.GET("/pong", func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{
					"messeage": "pong",
				})
			})
		}
		v2 := api.Group("v2")
		{
			v2.GET("/a", func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{
					"messeage": "a",
				})
			})

			v2.GET("/b", func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{
					"messeage": "b",
				})
			})
		}
	}

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

func ACustomMiddleware(context *gin.Context) {
	log.Println("I'm in a global middleware")
	context.Next()
}

func MyGroupV1Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("I'm in a group v1 middleware")
		if true {
			context.Next()
		}
	}
}
