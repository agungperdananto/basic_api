package main

import (
	"github.com/agungperdananto/basic_api/initializers"
	"github.com/agungperdananto/basic_api/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
