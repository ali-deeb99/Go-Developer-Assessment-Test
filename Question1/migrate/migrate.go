package main

import (
	"Question1/initializers"
	"Question1/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
}
