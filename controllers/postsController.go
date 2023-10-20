package controllers

import (
	"go-crud/initializers"
	"go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// Get data off req body

	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(201, gin.H{
		"post": result,
	})
}

func GetPosts(c *gin.Context) {

	var posts []models.Post

	// Get all records
	initializers.DB.Find(&posts)

	// Return it
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func GetPostByID(c *gin.Context) {
	// Get id off url
	id := c.Param("id")
	var post models.Post

	// get the post by ID
	initializers.DB.First(&post, id)

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})

}

func UpdatePostByID(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// get the data off req body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.First(&post, id)

	// update it

	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Return it
	c.JSON(204, gin.H{
		"post": post,
	})

}

func DeletePostByID(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	// get the post by ID
	initializers.DB.Delete(&models.Post{}, id)

	// Respond
	c.JSON(204, gin.H{
		"message": "Deleted post successfully",
	})
}
