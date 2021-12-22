package controllers

import (
	"clean_gin_api/lib"
	"clean_gin_api/models"
	"clean_gin_api/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserController data type
type SMSController struct {
	service services.SMSService
	logger  lib.Logger
}

// NewUserController creates new user controller
func NewSMSController(
	service services.SMSService,

	logger lib.Logger,
) SMSController {
	return SMSController{
		service: service,
		logger:  logger,
	}
}

// Save and Send message
func (s SMSController) SaveAndSendSms(c *gin.Context) {
	message := models.SMS{}
	if err := c.ShouldBindJSON(&message); err != nil {
		s.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := s.service.SendSMS(&message)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"err": err.Error(),
		})
	}
	fmt.Println(err, "----------")

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
	return
}
