package utils

import (
    "fmt"
    "strconv"
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
    smtp, err := ReadEnv("EMAIL_SMTP")
    if err != nil {
        return err
    }
    port, err := ReadEnv("EMAIL_PORT")
    if err != nil {
        return err
    }
    account, err := ReadEnv("EMAIL_ACCOUNT")
    if err != nil {
        return err
    }
    password, err := ReadEnv("EMAIL_PASSWORD")
    if err != nil {
        return err
    }
    portInt, err := strconv.Atoi(port)
    if err != nil {
        return fmt.Errorf("error converting port to int type in SendEmailVrfct: %v\n", err)
    }
    d := gomail.NewDialer(smtp, portInt, account, password)
    err = d.DialAndSend(m)
    return err
}
