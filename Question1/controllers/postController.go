package controllers

import (
	"Question1/initializers"
	"Question1/models"
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
)

// creating a user
func UserCreate(c *gin.Context) {

	var body struct {
		Name         string
		Phone_number string
	}

	var users []models.User
	initializers.DB.Find(&users)

	c.Bind(&body)

	// checking if the phone number existes in the data base

	var flag = false
	for i := 0; i < len(users); i++ {
		if users[i].Phone_number == body.Phone_number {
			flag = true
			break
		}
	}

	if flag == true {
		c.JSON(200, gin.H{
			"message": "The phone number must be unique",
		})
		return
	}

	user := models.User{Name: body.Name, Phone_number: body.Phone_number}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"user": "Creating your account successfully",
	})
}

func generateOTP() string {
	rand.Seed(time.Now().UnixNano())
	otp := fmt.Sprintf("%04d", rand.Intn(10000)) // Generate a 4-digit OTP
	return otp
}

// creating otp for the user
func GenerateOtp(c *gin.Context) {
	var body struct {
		Phone_number string
	}
	c.Bind(&body)
	var user models.User
	if err := initializers.DB.Where("Phone_number= ?", body.Phone_number).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "The phone number is wrong",
		})
		return
	}

	user.Otp = generateOTP()
	user.Otp_expiration_time = time.Now().Add(1 * time.Minute)
	initializers.DB.Save(&user)
	c.JSON(200, gin.H{
		"otp": user.Otp,
	})
}

// verify the phone number and the otp
func Verifyotp(c *gin.Context) {
	var body struct {
		Phone_number string
		Otp          string
	}
	c.Bind(&body)
	var user models.User
	if err := initializers.DB.Where("Phone_number= ? AND Otp= ?", body.Phone_number, body.Otp).First(&user).Error; err != nil {
		c.JSON(400, gin.H{
			"message": "credentials error",
		})
		return
	}
	if time.Now().After(user.Otp_expiration_time) {
		c.JSON(400, gin.H{
			"message": "The otp hsa expired",
		})
	} else {
		c.JSON(200, gin.H{
			"message": "Successfull",
		})
	}

}
