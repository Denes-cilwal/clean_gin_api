package routes

import (
	controllers "clean_gin_api/api/controllers"
	"clean_gin_api/infrastructure"
	"clean_gin_api/lib"
)

// SMSRoutes struct
type SMSRoutes struct {
	logger        lib.Logger
	handler       infrastructure.Router
	smsController controllers.SMSController
}

func NewSMSRoutes(
	logger lib.Logger,
	handler infrastructure.Router,
	smsController controllers.SMSController,

) SMSRoutes {
	return SMSRoutes{
		handler:       handler,
		logger:        logger,
		smsController: smsController,
	}
}

// Setup user routes
func (sr SMSRoutes) Setup() {
	sr.logger.Info("Setting up routes")
	api := sr.handler.Group("/api")
	{
		api.POST("/send-sms", sr.smsController.SaveAndSendSms)

	}
}
