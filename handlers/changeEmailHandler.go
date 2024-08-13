package handlers

import (
    "io"
    "errors"
    "fmt"
    "encoding/json"
    "net/http"
    "strings"
    "web-back-end/midware"
    "web-back-end/custypes"
    "web-back-end/database"
    "web-back-end/auth"
    "github.com/golang-jwt/jwt/v5"
)

func ChangeEmailHandler(w http.ResponseWriter, r *http.Request) {
    midware.SetCors(w)
    w.Header().Set("Content-Type", "application/json")

    //handle pre-flight request
    if strings.ToLower(r.Method) == "options" {
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

    //unmarshal received email
    type NewEmail struct {
        Email string 
    }
    var newEmail NewEmail
    err = json.Unmarshal(bodyBytes, &newEmail)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //declear response body
    var resBody custypes.ResponseBodyUser

    //get accessToken cookie and parse access token
    accessTokenCookie := auth.GetThisCookie("accessToken", r)
    
    if len(accessTokenCookie) != 0 {
        //extract access token
        accessTokenStr := strings.Split(accessTokenCookie, "=")[1]
        //parse access token
        accessToken, err := auth.ParseToken(accessTokenStr)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        //parse claims from access token
        claims, ok := accessToken.Claims.(jwt.MapClaims)
        if ok {
            userIdFloat, ok := claims["userId"].(float64)
            if ok {
                userId := int(userIdFloat) 
                //change email in database
                email, err := database.ChangeEmailById(userId, newEmail.Email, database.Blogdb) 
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                //change email_verified to false
                emailVerified, err := database.MarkEmailNotVerified(userId, database.Blogdb)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    fmt.Printf("failed to set email_verified to false, user id: %v\n", userId)
                    return
                }
                //assign new email to response body
                resBody.Email = email
                //assign emailVerified to EmailVerifid field in response body
                resBody.EmailVerified = emailVerified
            }else {
                http.Error(w, "error parsing user id from claims of access token", http.StatusInternalServerError)
                return
            }
        }else {
            http.Error(w, "error parsing claims from access token", http.StatusInternalServerError)
            return
        }
        //create and send back response body
        resBodyJson, err := json.Marshal(resBody)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        _, err = io.WriteString(w, string(resBodyJson))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        return
    }

    //get refresh token
    refreshTokenStr := auth.GetRefreshToken(r)
    //parse refresh token
    refreshToken, err := auth.ParseToken(refreshTokenStr)
    if err != nil {
        switch {
        case errors.Is(err, jwt.ErrTokenExpired):
            http.Error(w, err.Error(), http.StatusUnauthorized)
            return
        case errors.Is(err, jwt.ErrTokenMalformed):
            http.Error(w, err.Error(), http.StatusUnauthorized)
            return
        default:
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
    //extract claims from refresh token
    claims, ok := refreshToken.Claims.(jwt.MapClaims)
    if ok {
        userIdFloat, ok := claims["userId"].(float64)
        if ok {
            userId := int(userIdFloat)
            //change email in database
            email, err := database.ChangeEmailById(userId, newEmail.Email, database.Blogdb)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            //change email_verified to false
            emailVerified, err := database.MarkEmailNotVerified(userId, database.Blogdb)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                fmt.Printf("failed to set email_verified to false, user id: %v\n", userId)
                return
            }
            //create and send back response body
            resBody.Email = email
            resBody.EmailVerified = emailVerified
            resBodyJson, err := json.Marshal(resBody)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            _, err = io.WriteString(w, string(resBodyJson))
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
        }else {
            http.Error(w, "error parsing user id from claims of refresh token", http.StatusInternalServerError)
            return
        }
    }else {
        http.Error(w, "error parsing claims from refresh token", http.StatusInternalServerError)
        return
    }

}
