package services

import (
	"clean_gin_api/lib"

	"github.com/twilio/twilio-go"
)

type TwilioSmsService struct {
	client *twilio.RestClient
	logger lib.Logger
	env    lib.Env
}

func NewSmsService(
	env lib.Env,
	logger lib.Logger,
) TwilioSmsService {

	accountid := env.AccountSid
	authToken := env.AuthToken
	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username: accountid,
		Password: authToken,
	})
	return TwilioSmsService{
		client: client,
		logger: logger,
	}
}

type SMSParams struct {
	To   string
	From string
	Body string
}
