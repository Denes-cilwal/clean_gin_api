package controllers

import (
	"clean_gin_api/models"
	"clean_gin_api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	Userservice services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return UserController{
		Userservice: userService,
	}
}

// SaveUser saves the user
func (uc UserController) SaveUser(c *gin.Context) {
	// validate input
	user := models.User{}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if err := uc.Userservice.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"data": "user created"})
}
