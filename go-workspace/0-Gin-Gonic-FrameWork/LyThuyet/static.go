package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static_file", "./assets")
	r.Use(MyACustomMiddleware())

	r.GET("/ping", getPing)
	r.POST("/ping", postPing)
	r.GET("/detail/:id", getDetail)

	//upload single file
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", func(context *gin.Context) {
		// single file
		file, _ := context.FormFile("file")
		log.Println(file.Filename)

		// Upload the file to specific dst.
		context.SaveUploadedFile(file, "./assets/upload/"+file.Filename)

		context.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	//upload multiple file
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload_multiple_file", func(context *gin.Context) {
		// Multipart form
		form, _ := context.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// Upload the file to specific dst.
			context.SaveUploadedFile(file, "./assets/upload/"+uuid.New().String()+filepath.Ext(file.Filename))
		}
		context.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	api := r.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			//Khởi tạo middleware chỉ dùng cho group có prefix là v1
			v1.Use(MyGroupV1Middleware())

			v1.GET("/ping", func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{
					"message": "Ping",
				})
			})

			v1.GET("/pong", func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{
					"message": "pong",
				})
			})
		}

		v2 := api.Group("/v2")
		{
			v2.GET("/a", func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{
					"message": "a",
				})
			})
			v2.GET("/b", func(context *gin.Context) {
				context.JSON(http.StatusOK, gin.H{
					"message": "b",
				})
			})
		}
	}

	//api/v1/ping
	//api/v1/pong

	//api/v2/a
	//api/v2/b

	r.Run(":333")
}

func getDetail(context *gin.Context) {
	log.Println("I'm in getDetail handler")
	id := context.Param("id")

	context.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

func getPing(context *gin.Context) {
	log.Println("I'm in getPing handler")
	name := context.DefaultQuery("name", "Guest")

	var data = map[string]interface{}{
		"message": "Hello " + name + " from GET ping",
	}

	//gin.H{
	//			"message" : "Hi from ping",
	//}

	context.JSON(http.StatusOK, data)
}

func postPing(context *gin.Context) {
	log.Println("I'm in postPing handler")
	address := context.DefaultPostForm("addr", "Viet Nam")

	context.JSON(http.StatusOK, gin.H{
		"message": "Hello from " + address + " POST ping",
	})
}

func MyGroupV1Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("I'm in a group v1 middleware")
		context.Next()
	}
}

func MyACustomMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Println("I'm in a global middleware")
		if true {
			context.Next()
		}
	}
}

//func MyACustomMiddleware(context *gin.Context)  {
//	log.Println("I'm in a global middleware")
//	if(true){
//		context.Next()
//	}
//}
