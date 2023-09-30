package main

import (
	"encoding/json"
	"net/http"

	"github.com/agungperdananto/basic_api/controllers"
	"github.com/agungperdananto/basic_api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/posts", controllers.PostsCreate)
	r.PUT("/posts/:id", controllers.PostUpdate)

	r.GET("/posts", controllers.PostsIndex)
	r.GET("/posts/:id", controllers.PostShow)

	r.DELETE("/posts/:id", controllers.PostDelete)

	// Test request
	r.GET("/direct", func(c *gin.Context) {
		response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
		if err != nil {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		// if there was no error, you should close the body
		defer response.Body.Close()

		// hence this condition is moved into its own block
		if response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}

		// use a proper struct type in your real code
		// the empty interface is just for demonstration
		var v interface{}
		json.NewDecoder(response.Body).Decode(&v)

		c.JSON(200, v)
	})
	r.Run()
}
