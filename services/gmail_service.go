package services

import (
	"clean_gin_api/lib"
	"clean_gin_api/utils"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// GmailService -> GmailService structure
type GmailService struct {
	gmailService *gmail.Service
	logger       lib.Logger
	env          lib.Env
}

// NewGmailService -> Constructor Function
func NewGmailService(
	env lib.Env,
	logger lib.Logger,
) GmailService {
	ctx := context.Background()

	oauthConfig := oauth2.Config{
		ClientID:     env.MailClientID,
		ClientSecret: env.MailClientSecret,
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost",
		Scopes:       []string{"https://www.googleapis.com/auth/gmail.send"},
	}

	expiry, _ := time.Parse("2006-01-02", "2018-04-16")
	token := oauth2.Token{
		RefreshToken: env.MailTokenType,
		TokenType:    "Bearer",
		Expiry:       expiry,
	}

	var tokenSource = oauthConfig.TokenSource(ctx, &token)

	srv, err := gmail.NewService(ctx, option.WithTokenSource(tokenSource))

	if err != nil {
		logger.Fatal("failed to retrieve gmail client", err.Error())
	}

	logger.Info("gmail client created successfully")

	return GmailService{
		gmailService: srv,
		logger:       logger,
	}
}

type EmailParams struct {
	To              string
	Subject         string
	SubjectData     interface{}
	SubjectTemplate string
	Body            string
	BodyData        interface{}
	BodyTemplate    string
	Lang            string
}

// Send Email -> sends email with provided data
func (g GmailService) SendEmail(params EmailParams) (bool, error) {
	isEmailSubjectEmpty := utils.IsInterfaceEmpty(params.SubjectData)
	var err error

	if !isEmailSubjectEmpty {
		params.Subject, err = utils.ParseTemplate(params.SubjectTemplate, params.SubjectData)
		if err != nil {
			return false, errors.New("unable to parse email subject template")
		}
	}

	emailBody, err := utils.ParseTemplate(params.BodyTemplate, params.BodyData)
	if err != nil {
		return false, errors.New("unable to parse email body template")
	}

	var message gmail.Message

	emailTo := "To: " + params.To + "\r\n"
	subject := "Subject: " + params.Subject + "\r\n"

	msgStr := emailTo + subject + "\n" + emailBody
	var msg []byte

	if params.Lang != "en" {
		msgStrJP, _ := utils.ToISO2022JP(msgStr)
		msg = msgStrJP
	} else {
		msg = []byte(msgStr)
	}
	message.Raw = base64.URLEncoding.EncodeToString(msg)

	// Send the message
	_, err = g.gmailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Printf("Error: %v", err)
		return false, err
	}

	fmt.Println("Message sent!")
	return true, nil
}
