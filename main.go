package gmailer

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"

	"gitlab.com/hartsfield/goken"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

// Message allows us to send messages
type Message struct {
	Recipient string
	Subject   string
	Body      string
}

// getClient uses a Context and Config to retrieve a Token
// then generate a Client. It returns the generated Client.
func getClient(ctx context.Context) *http.Client {
	b, err := ioutil.ReadFile("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	/////////////////////////////////////////////////////////////////////////////
	// SCOPE https://developers.google.com/gmail/api/auth/scopes
	/////////////////////////////////////////////////////////////////////////////
	// If modifying these scopes, delete your previously saved credentials
	// at ~/.credentials/gmail-go-quickstart.json
	// If modifying this script to access other features of gmail, you may need
	// to change scope.
	config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	tok := goken.NewTokenizer()
	cacheFile, err := tok.TokenCacheFile()
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	token, err := tok.TokenFromFile(cacheFile)
	if err != nil {
		token = tok.GetTokenFromWeb(config)
		tok.SaveToken(cacheFile, token)
	}
	return config.Client(ctx, token)
}

// Send allows us to send the email
func (m *Message) Send() {
	ctx := context.Background()
	client := getClient(ctx)

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve gmail Client %v", err)
	}

	// Make and send the message
	var message gmail.Message

	messageStr := []byte(
		"From: xxx\r\n" +
			"To: " + m.Recipient + "\r\n" +
			"Subject: " + m.Subject + "\r\n\r\n" +
			m.Body)

	message.Raw = base64.URLEncoding.EncodeToString(messageStr)

	_, err = srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		log.Print("Email sent")
	}

}
