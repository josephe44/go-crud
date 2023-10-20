package main

import (
	"go-crud/initializers"
	"go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	// Migrate the schema
	initializers.DB.AutoMigrate(&models.Post{})
}
