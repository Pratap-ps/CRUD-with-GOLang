package controlers

import (
	"examples/go-crud/initializers"
	"examples/go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post) // pass pointer of data to Create

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Post": post,
	})
}

func PostIndex(c *gin.Context) {
	// Get posts
	var posts []models.Post
	initializers.DB.Find(&posts)
	// Respond with posts
	c.JSON(200, gin.H{
		"Posts": posts,
	})
}

func PostSingle(c *gin.Context) {
	// Get id of url
	id := c.Param("id")
	// Get posts
	var post models.Post
	initializers.DB.First(&post, id)
	// Respond with posts
	c.JSON(200, gin.H{
		"Post": post,
	})
}

func PostUpdate(c *gin.Context) {
	// Get the id of the URL
	id := c.Param("id")
	// Get the data of req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	// Find the post for updating
	var post models.Post
	initializers.DB.First(&post, id)
	// Update the post
	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Body: body.Title})
	// Respond with post
	c.JSON(200, gin.H{
		"Post": post,
	})

}

func PostDelete(c *gin.Context) {
	// Get the id of the URL
	id := c.Param("id")
	// Delete the post
	initializers.DB.Delete(&models.Post{}, id)
	// Respond with post
	c.Status(200)

}
