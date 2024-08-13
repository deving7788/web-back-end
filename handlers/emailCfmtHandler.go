package handlers

import (
    "fmt"
    "strings"
    "net/http"
    "time"
    "database/sql"
    "strconv"
    "web-back-end/midware"
    "web-back-end/database"
)

func EmailCfmtHandler(w http.ResponseWriter, r *http.Request) {
    midware.SetCors(w)
    w.Header().Set("Content-Type", "application/json")
    //handle pre-flight request
    if strings.ToLower(r.Method) == "options" {
        return
    }
    //get id and verification token from request url
    url := *r.URL
    rawQuery := url.RawQuery
    qStrs := strings.Split(rawQuery, "&")

    vrfctToken := strings.Split(qStrs[0], "=")[1] 
    tokenIdStr := strings.Split(qStrs[1], "=")[1]
    tokenId, err := strconv.Atoi(tokenIdStr)
    if err != nil {
        http.Error(w, "error converting string", http.StatusInternalServerError)
        return
    }
    //retrieve verification token and expiry time from database
    vrfctTokenDb, expiryTime, err := database.GetVrfctTokenStrAndTime(tokenId, database.Blogdb)
    if err != nil {
        switch {
        case err == sql.ErrNoRows:
            fmt.Println("no rows error")
            return
        default:
            fmt.Println("some error:", err)
            return
        }
    }else {
        if vrfctTokenDb == vrfctToken {
            timeNow := time.Now()
            if timeNow.After(expiryTime) {
                fmt.Println("The verification token has expired.")
                return
            }else {
                err = database.MarkEmailVerified(tokenId, database.Blogdb) 
                if err != nil {
                    switch {
                    case nil == sql.ErrNoRows:
                        fmt.Println("no rows error")
                        return
                    default:
                        fmt.Println("some error:", err)
                        return
                    }
                }else {
                    fmt.Println("This email has been verified")
                    return
                }
            }
            fmt.Println("they are equal")
        }else {
            fmt.Println("verification tokens are different.")
            return
        }
    }

}
