package utils

import (
    "fmt"
    "os"
    "strconv"
    "gopkg.in/gomail.v2"
)

func SendEmailVrfctEmail(displayName string, email string, vrfctLinkStr string) error {

    body := fmt.Sprintf( `<div style="font-size: 22px; font-weight: 600; color: #114211;">
                            Hello %s,
                          </div>
                          <br/>
                          <div style="font-size: 20px; font-weight: 600; color: #114211;">Thank you for joining.</div>
                          <div style="font-size: 20px; font-weight: 600; color: #114211;">
                            Please click <a href=%s>here</a> to activate your account.
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
                          `, displayName, vrfctLinkStr)
    m := gomail.NewMessage()
    m.SetHeader("From", "quarque@quarque.com")
    m.SetHeader("To", fmt.Sprintf("%s", email))
    m.SetHeader("Subject", "email verification")
    m.SetBody("text/html", body)
    smtp := os.Getenv("EMAIL_SMTP")
    port := os.Getenv("EMAIL_PORT")
    account := os.Getenv("EMAIL_ACCOUNT")
    password := os.Getenv("EMAIL_PASSWORD")
    portInt, err := strconv.Atoi(port)
    if err != nil {
        return fmt.Errorf("error converting port to int type: %v\n", err)
    }
    d := gomail.NewDialer(smtp, portInt, account, password)
    err = d.DialAndSend(m)
    return err
}
