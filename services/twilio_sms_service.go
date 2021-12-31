package services

import (
	"clean_gin_api/lib"
	"clean_gin_api/models"
	"clean_gin_api/respository"
	"encoding/json"
	"fmt"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

type SMSService struct {
	repository respository.SMSRepository
	logger     lib.Logger
	env        lib.Env
}

// NewSMSService is a factory function for
// initializing a SMSService with its repository layer dependencies
func NewSMSService(smsrepo respository.SMSRepository, logger lib.Logger, env lib.Env) SMSService {
	return SMSService{
		repository: smsrepo,
		logger:     logger,
		env:        env,
	}
}

func (ss SMSService) SendSMS(sms *models.SMS) error {

	// var clientMsg models.SMS
	sid := ss.env.AccountSid
	authToken := ss.env.AuthToken
	senderNumber := ss.env.SmsSenderNumber

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: sid,
		Password: authToken,
	})
	params := &openapi.CreateMessageParams{}
	params.SetTo(sms.To)
	params.SetFrom(senderNumber)
	params.SetBody(sms.Message)

	resp, err := client.ApiV2010.CreateMessage(params)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}

	// err := ss.repository.SendSMS(response)
	return nil

}
