 package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"api/constants"
	

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//SendGridEmailService -> Interface to Sendgrid Email Service

type SendGridEmailService interface {
SendEmail(toEmail, emailSubject string,  emailBody string,  lang string, isBCC bool) (bool, error)
	GetAllUnsubscribers() ([]string, error)
	}

type sendGridEmailService struct {
	SendGridEmailService *sendgrid.Client
}

// NewSendGridEmailService : initialize
func NewSendGridEmailService() SendGridEmailService {

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))

	return &sendGridEmailService{
		SendGridEmailService: client,
	}
}

//SendEmail -> to send email with the provided data
func (send *sendGridEmailService) SendEmail(toEmail, emailSubject string,  emailBody string,  lang string, isBCC bool) (bool, error) {

	
  // Add from and to address
	
	from := mail.NewEmail("e運営チーム", os.Getenv("FROM_EMAIL"))
	subject := emailSubject

	to := mail.NewEmail("", toEmail)

  // define content type here
	content := mail.NewContent("text/html", emailBody)

  // create new *Personalization
	personalization := mail.NewPersonalization()

  // populate `personalization` with data
	if isBCC {
		bcc := mail.NewEmail("", os.Getenv("BCC_EMAIL"))
		personalization.AddBCCs(bcc)
	}

	personalization.AddTos(to)

	message := mail.NewV3MailInit(from, subject, to, content)

	message.AddPersonalizations(personalization)

  // Send the message using sendgrid service

	emailRes, err := send.SendGridEmailService.Send(message)
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		fmt.Println("Message sent!", emailRes)
	}
	return true, nil
}

func (send *sendGridEmailService) GetAllUnsubscribers() ([]string, error) {

	client := os.Getenv("SENDGRID_API_KEY")
	host := constants.Host
	request := sendgrid.GetRequest(client, constants.UnsubscribesEndpoint, host)
	request.Method = "GET"
	response, err := sendgrid.API(request)
	if err != nil {
		return nil, err
	}
	var unsubscribedEmails []string

	var dat []map[string]interface{}

	if err := json.Unmarshal([]byte(response.Body), &dat); err != nil {
		return nil, err
	}

	for _, res := range dat {
		unsubscribedEmails = append(unsubscribedEmails, res["email"].(string))
	}
	return unsubscribedEmails, nil

}
