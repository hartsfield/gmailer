// Package gmailer is a go library used for sending emails using the gmail API.

// Copyright (c) 2017 J. Hartsfield

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
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
