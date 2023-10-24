package main

import (
	"go-crud/controllers"
	"go-crud/initializers"
	"go-crud/middleware"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Posts
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.UpdatePostByID)
	r.DELETE("/posts/:id", controllers.DeletePostByID)

	r.GET("/", controllers.GetPosts)
	r.GET("/posts/:id", controllers.GetPostByID)

	// Users
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.POST("/logout", controllers.Logout)

	r.GET("/me", middleware.AuthMiddleware, controllers.Me)

	r.Run()
}
