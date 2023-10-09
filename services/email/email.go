package email

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/mailgun/mailgun-go/v4"
)

// Your available domain names can be found here:
// (https://app.mailgun.com/app/domains)
var yourDomain string = os.Getenv("MAILGUN_DOMAIN") // e.g. mg.yourcompany.com

// You can find the Private API Key in your Account Menu, under "Settings":
// (https://app.mailgun.com/app/account/security)
var privateAPIKey string = os.Getenv("MAILGUN_API_KEY")

func Send(subject, body, recipient string) (string, string, error) {
	// Create an instance of the Mailgun Client
	mg := mailgun.NewMailgun(yourDomain, privateAPIKey)

	//When you have an EU-domain, you must specify the endpoint:
	//mg.SetAPIBase("https://api.eu.mailgun.net/v3")

	sender := os.Getenv("MAILGUN_SENDER")

	// The message object allows you to add attachments and Bcc recipients
	message := mg.NewMessage(sender, subject, body, recipient)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if privateAPIKey == "LOCALHOST" {
		id := "dummy id"
		resp := "mock sent"
		log.Printf("mockemail: to:%s subject:%s body:%s", recipient, subject, body)
		return id, resp, nil
	}

	// Send the message with a 10 second timeout
	resp, id, err := mg.Send(ctx, message)

	if err != nil {
		return "", "", err
	}
	log.Printf("sent email: %s %s %s; %s %s", subject, body, recipient, resp, id)

	return id, resp, nil
}
