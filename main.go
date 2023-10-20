package main

import (
	"go-crud/controllers"
	"go-crud/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.UpdatePostByID)
	r.DELETE("/posts/:id", controllers.DeletePostByID)

	r.GET("/", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostByID)

	r.Run()
}
