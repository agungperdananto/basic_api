package controllers

import (
	"errors"

	"github.com/agungperdananto/basic_api/initializers"
	"github.com/agungperdananto/basic_api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post
	// get the post
	initializers.DB.Find(&posts)
	// respond them
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	// get id off url
	id := c.Param("id")
	var post models.Post
	// get the post
	result := initializers.DB.First(&post, id)
	// check error
	// result.RowsAffected // returns count of records found
	// result.Error        // returns error or nil

	// check error ErrRecordNotFound
	err := errors.Is(result.Error, gorm.ErrRecordNotFound)
	if err {
		c.JSON(400, gin.H{
			"post": result.Error.Error(),
		})
		return
	}
	// respond them
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	//  get the id
	id := c.Param("id")
	// get the data req body
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)
	// find the post were updating
	var post models.Post
	result := initializers.DB.First(&post, id)
	err := errors.Is(result.Error, gorm.ErrRecordNotFound)
	if err {
		c.JSON(400, gin.H{
			"post": result.Error.Error(),
		})
		return
	}

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Body:  body.Body,
		Title: body.Title,
	})
	// Respond with it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	// get id off url
	id := c.Param("id")
	// delete the post
	initializers.DB.Delete(&models.Post{}, id)
	// respond them
	c.Status(200)
}
