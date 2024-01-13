package main

import (
	"Question1/controllers"
	"Question1/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}
func main() {

	r := gin.Default()
	r.POST("/POST/api/users", controllers.UserCreate)
	r.POST("/POST/api/users/generateotp", controllers.GenerateOtp)
	r.POST("/POST/api/users/verifyotp", controllers.Verifyotp)

	r.Run()
}
