package controllers

import (
	"github.com/agungperdananto/basic_api/initializers"
	"github.com/agungperdananto/basic_api/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data off req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// create post
	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}
