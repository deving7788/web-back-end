package utils

import (
    "fmt"
    "strconv"
    "gopkg.in/gomail.v2"
)

func SendPasswordResetEmail(email string, prLinkStr string) error {

    body := fmt.Sprintf( `<div style="font-size: 20px; font-weight: 600; color: #114211;">
                            Please click <a href=%s>here</a> to reset your password.
                          </div>
                          <br/>
                          <br/>
                          <div style="font-size: 16px; color: #114211;">from quarque.com</div>
                          <br/>
                          <br/>
                          <br/>
                          <br/>
                          <br/>
                          <br/>
                          <div style="font-size: 14px; color: #114211;">If you did not request this email, please ignore it.</div>
                          `, prLinkStr)
    m := gomail.NewMessage()
    m.SetHeader("From", "quarque@quarque.com")
    m.SetHeader("To", fmt.Sprintf("%s", email))
    m.SetHeader("Subject", "password reset")
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
        return fmt.Errorf("error converting port to int type: %v\n", err)
    }
    d := gomail.NewDialer(smtp, portInt, account, password)
    err = d.DialAndSend(m)
    return err
}
