package handlers

import (
    "os"
    "time"
    "io"
    "net/http"
    "strings"
    "strconv" 
    "database/sql"
    "web-back-end/midware"
    "web-back-end/database"
)

func SendPasswordResetPageHandler(w http.ResponseWriter, r *http.Request) {
    midware.SetCors(w)
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    //handle pre-flight request
    if strings.ToLower(r.Method) == "options" {
        return
    }
    //get pr token id and pr token from request url
    prToken := r.URL.Query()["token"][0]
    prTokenIdStr := r.URL.Query()["id"][0] 
    prTokenId, err := strconv.Atoi(prTokenIdStr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //declear response of type []byte
    var resBytes []byte
    //retrieve password reset token and expiry time from database, and respond
    prTokenDb, expiryTime, err := database.GetPrTokenStrAndTime(prTokenId, database.Blogdb)
    if err != nil {
        switch {
        case err == sql.ErrNoRows:
            resBytes, err = os.ReadFile("htmls/passwordReset404.html")
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            _, err = io.WriteString(w, string(resBytes))
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            return
        default:
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }else {
        if prTokenDb == prToken {
            timeNow := time.Now()
            if timeNow.After(expiryTime) {
                resBytes, err = os.ReadFile("htmls/passwordResetTokenExpired.html")
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                _, err = io.WriteString(w, string(resBytes))
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                return
            }else {
                resBytes, err = os.ReadFile("htmls/passwordReset.html")
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                _, err = io.WriteString(w, string(resBytes))
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                return
            }
        }else {
            resBytes, err = os.ReadFile("htmls/passwordResetRequestInvalid.html")
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            _, err := io.WriteString(w, string(resBytes))
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            return
        }
    }
}
