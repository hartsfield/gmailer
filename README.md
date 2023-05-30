# GMAILER

A tiny library for sending gmails. I use it in my programs to send me email 
alerts when things happen that I need to know about. Simply form a 
`gmailer.Message` and then call its method `Send()`

NOTE:

On the first run of this app it will look for a `client_secret.json` file in 
the root directory of your project. It will then prompt you to visit a page to
copy/paste an access token from the resulting web address back into the 
terminal. It will then store a file in `~/.credentials` that will be used every
time to conect to the gmail api from then on.

`client_secret.json` can be generated from within an app on google cloud 
platform.

If other programs attempt to interact with the gmail api, one of these files 
could cause an error. 

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
