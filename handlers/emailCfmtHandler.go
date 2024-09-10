package handlers

import (
    "io"
    "os"
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
    w.Header().Set("Content-Type", "text/html")
    //handle pre-flight request
    if strings.ToLower(r.Method) == "options" {
        return
    }
    //get id and verification token from request url
    url := r.URL
    rawQuery := url.RawQuery
    qStrs := strings.Split(rawQuery, "&")

    vrfctToken := strings.Split(qStrs[0], "=")[1] 
    tokenIdStr := strings.Split(qStrs[1], "=")[1]
    tokenId, err := strconv.Atoi(tokenIdStr)
    if err != nil {
        http.Error(w, "error converting string", http.StatusInternalServerError)
        return
    }
    //declear response of type []byte
    var resBytes []byte
    //retrieve verification token and expiry time from database
    vrfctTokenDb, expiryTime, err := database.GetVrfctTokenStrAndTime(tokenId, database.Blogdb)
    if err != nil {
        switch {
        case err == sql.ErrNoRows:
            resBytes, err = os.ReadFile("handlers/../htmls/emailVrfct404.html")
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
        if vrfctTokenDb == vrfctToken {
            timeNow := time.Now()
            if timeNow.After(expiryTime) {
                resBytes, err = os.ReadFile("handlers/../htmls/emailVrfctTokenExpired.html")
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
                err = database.MarkEmailVerified(tokenId, database.Blogdb) 
                if err != nil {
                    switch {
                    case nil == sql.ErrNoRows:
                        resBytes, err = os.ReadFile("handlers/../htmls/emailVrfct404.html")
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
                    resBytes, err = os.ReadFile("handlers/../htmls/emailVerified.html")
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
            }
        }else {
            resBytes, err = os.ReadFile("handlers/../htmls/emailVrfctRequestInvalid.html")  
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
