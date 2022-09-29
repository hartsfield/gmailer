# GMAILER

A tiny library for sending gmails. I use it in my programs to send me email 
alerts when things happen that I need to know about. Simply form a 
`gmailer.Message` and then call its method `Send()`

Example:

        package main

        import (
                "fmt"
                "github.com/hartsfield/gmailer"
        )

        func main() {
                msg := gmailer.Message{
                        Recipient: "YOUR___GMAIL___HERE@gmail.com",
                        Subject:   "ALERT! Server has exceeded the maximum number of open file descriptors",
                        Body:      "...",
                }
                msg.Send(onMessageSent)
        }

        func onMessageSent() {
                fmt.Println("Message sent successfully")
        }
