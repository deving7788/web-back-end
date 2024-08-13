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
    m.SetHeader("From", "frank@gobackend.com")
    m.SetHeader("To", fmt.Sprintf("%s", email))
    m.SetHeader("Subject", "greetings")
    m.SetBody("text/html", body)
    d := gomail.NewDialer("sandbox.smtp.mailtrap.io", 587, "779e18e63fb8db", "cc3e4f50ee9b01")
    err := d.DialAndSend(m)
    return err
}
