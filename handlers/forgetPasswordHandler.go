package handlers

import (
    "fmt"
    "net/http"
    "strings"
    "web-back-end/midware"
)

func ForgetPasswordHandler(w http.ResponseWriter, r *http.Request) {
    midware.SetCors(w)
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    //hanlde pre-fight request
    if strings.ToLower(r.Method) == "options" {
        return
    }
    //get email address from request url
    email := r.URL.Query()["email"][0]
    fmt.Printf("email in ForgetPasswordHandler: %s", email+"haha")
}
