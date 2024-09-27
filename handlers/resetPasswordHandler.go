package handlers

import (
    "net/http"
    "time"
    "errors"
    "strconv"
    "io"
    "os"
    "database/sql"
    "encoding/json"
    "web-back-end/midware"
    "web-back-end/database"
    "web-back-end/utils"
)

func ResetPasswordHandler(w http.ResponseWriter, r *http.Request) {
    midware.SetCors(w)
    w.Header().Set("Content-Type", "application/json")
    //handle pre-flight request
    if r.Method == "options" {
        return
    }
    //get request body of json
    bodyStream := r.Body
    bodyBytes, err := io.ReadAll(bodyStream)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer bodyStream.Close()
    //unmarshal received content
    type PrBody struct {
        Password string `json:"password,omitempty"`
        PrToken string `json:"token,omitempty"`
        PrTokenId int `json:"id,omitempty"`
    }
    var prBody PrBody
    err = json.Unmarshal(bodyBytes, &prBody)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //retrieve password reset token and expiry time from database, and respond
    prTokenDb, expiryTime, err := database.GetPrTokenStrAndTime(prBody.PrTokenId, database.Blogdb)
    if err != nil {
        switch {
        case err == sql.ErrNoRows:
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        default:
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }else {
        if prTokenDb == prBody.PrToken {
            timeNow := time.Now()
            if timeNow.After(expiryTime) {
                http.Error(w, "time out", http.StatusRequestTimeout)
                return
            }else {
                userId, err := database.GetIdByPrTokenId(prBody.PrTokenId, database.Blogdb)
                if err != nil {
                    switch {
                    case errors.Is(err, sql.ErrNoRows):
                        http.Error(w, err.Error(), http.StatusNotFound)
                        return
                    default:
                        http.Error(w, err.Error(), http.StatusInternalServerError)
                        return
                    }
                }
                //hash and store password
                costStr := os.Getenv("COST")
                cost, err := strconv.Atoi(costStr)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                hashedPasswordStr, err := utils.HashPassword(prBody.Password, cost)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                err = database.UpdatePasswordById(userId, hashedPasswordStr, database.Blogdb)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                w.WriteHeader(200)
                return
            }
        }else {
            http.Error(w, "invalid", http.StatusBadRequest)
            return
        }
    }
}
