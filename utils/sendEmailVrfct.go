package utils

import (
    "fmt"
    "gopkg.in/gomail.v2"
)

func SendEmailVrfct(displayName string, email string, vrfctLinkStr string) error {

    body := fmt.Sprintf( `<div>
                            Hello <b>%s</b>
                          </div>
                          <div style="font-size: 20px; background-color: green;">
                            click <a href=%s>here</a> to activate your account
                          </div>`, displayName, vrfctLinkStr)
    m := gomail.NewMessage()
    m.SetHeader("From", "quarque@quarque.com")
    m.SetHeader("To", fmt.Sprintf("%s", email))
    m.SetHeader("Subject", "greetings")
    m.SetBody("text/html", body)
    err := d.DialAndSend(m)
    return err
}
