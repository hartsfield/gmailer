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
	"log"

	"gitlab.com/hartsfield/gmailAPI"
	"golang.org/x/net/context"
	"google.golang.org/api/gmail/v1"
)

// Message allows us to send messages
type Message struct {
	Recipient string
	Subject   string
	Body      string
}

// Send allows us to send the email
func (m *Message) Send(cb func()) {
	// Connect to the gmail API service.
	ctx := context.Background()
	srv := gmailAPI.ConnectToService(ctx, gmail.GmailSendScope)

	// Make and send the message.
	var message gmail.Message

	messageStr := []byte(
		"From: xxx\r\n" +
			"To: " + m.Recipient + "\r\n" +
			"Subject: " + m.Subject + "\r\n\r\n" +
			m.Body)

	message.Raw = base64.URLEncoding.EncodeToString(messageStr)

	_, err := srv.Users.Messages.Send("me", &message).Do()
	if err != nil {
		log.Printf("Error: %v", err)
	} else {
		cb()
	}

}
