package controllers

import (
	"go-crud/initializers"
	"go-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserStruct struct {
	ID    int
	Email string
	// Add more fields as needed
}

func PostsCreate(c *gin.Context) {

	// Get data off req body

	var body struct {
		Title      string
		Body       string
		PostUserID uint
	}
	c.Bind(&body)

	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user request"})
		return
	}

	// Create a post
	post := models.Post{Title: body.Title, Body: body.Body, PostUserID: userID.(uint)}

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

	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user request"})
		return
	}

	var posts []models.Post

	// Get all matched records
	initializers.DB.Where("post_user_id = ?", userID).Find(&posts)

	// initializers.DB.Find(&posts)

	// Return it
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func GetPostByID(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user request"})
		return
	}

	var post models.Post

	// get the post by ID
	initializers.DB.Where("post_user_id = ?", userID).First(&post, id)

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})

}

func UpdatePostByID(c *gin.Context) {
	// Get id off url
	id := c.Param("id")

	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user request"})
		return
	}

	// get the data off req body
	var body struct {
		Title string
		Body  string
	}
	c.Bind(&body)

	// Find the post were updating
	var post models.Post
	initializers.DB.Where("post_user_id = ?", userID).First(&post, id)

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

	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user request"})
		return
	}

	// get the post by ID
	initializers.DB.Where("post_user_id = ?", userID).Delete(&models.Post{}, id)

	// Respond
	c.JSON(204, gin.H{
		"message": "Deleted post successfully",
	})
}
