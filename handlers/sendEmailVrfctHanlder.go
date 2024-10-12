package handlers

import (
    "strings"
    "fmt"
    "net/http"
    "errors"
    "os"
    "web-back-end/midware"
    "web-back-end/auth"
    "web-back-end/utils"
    "web-back-end/database"
    "github.com/golang-jwt/jwt/v5"
)

func SendEmailVrfctHandler(w http.ResponseWriter, r *http.Request) {
    midware.SetCors(w)
    w.Header().Set("Content-Type", "application/json")
    //handle pre-flight request
    if strings.ToLower(r.Method) == "options" {
        return
    }
    //declare and initialze variables
    var displayName, email string = "", ""
    var userId int
    var userIdFloat float64
    apiAddress := os.Getenv("API_HOST")
    baseStr := fmt.Sprintf("%s/api/user/email-cfmt?", apiAddress)

    //get accessToken cookie and parse access token
    accessTokenStr := auth.GetThisCookie("accessToken", r)
    if len(accessTokenStr) != 0 {
        //parse access token
        accessToken, err := auth.ParseToken(accessTokenStr)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        //parse claims from accessToken
        claims, ok := accessToken.Claims.(jwt.MapClaims)
        if ok {
            //parse user id
            userIdFloat, ok = claims["userId"].(float64)
            if !ok {
                http.Error(w, "error parsing user id from access token", http.StatusInternalServerError)
                return
            }
            userId = int(userIdFloat)
            //parse display name
            displayName, ok = claims["displayName"].(string)
            if !ok {
                http.Error(w, "error parsing displayName from access token", http.StatusInternalServerError)
                return
            }
            //parse email
            email, ok = claims["email"].(string)
            if !ok {
                http.Error(w, "error parsing email from access token", http.StatusInternalServerError)
                return
            }
            //generate an array of 30 random bytes 
            vrfctTokenBytes := utils.GenerateRandomBytes(30);
            vrfctTokenId, err := database.StoreEmailVrfctToken(vrfctTokenBytes, userId, database.Blogdb)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            //generate a verification link
            vrfctLinkStr := baseStr + "token=" + string(vrfctTokenBytes) + "&id=" + fmt.Sprintf("%d", vrfctTokenId) 
            //send email verification email
            err = utils.SendEmailVrfctEmail(displayName, email, vrfctLinkStr)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            return

        }else {
            http.Error(w, "error parsing claims from access token", http.StatusInternalServerError)
            return
        }
    }

    //get refresh token
    refreshTokenStr := auth.GetRefreshToken(r)
    //parse refresh token
    refreshToken, err := auth.ParseToken(refreshTokenStr)
    //handle parsing error and expired token
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
        //parse user id
        userIdFloat, ok = claims["userId"].(float64)
        if !ok {
            http.Error(w, "error parsing user id from refresh token", http.StatusInternalServerError)
            return
        }
        userId = int(userIdFloat)
        //parse display name
        displayName, ok = claims["displayName"].(string)
        if !ok {
            http.Error(w, "error parsing displayName from refresh token", http.StatusInternalServerError)
            return
        }
        //parse email
        email, ok = claims["email"].(string)
        if !ok {
            http.Error(w, "error parsing email from refresh token", http.StatusInternalServerError)
            return
        }
        //generate an array of 30 random bytes
        vrfctTokenBytes := utils.GenerateRandomBytes(30);
        vrfctTokenId, err := database.StoreEmailVrfctToken(vrfctTokenBytes, userId, database.Blogdb)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        //generate a verification link
        vrfctLinkStr := baseStr + "token=" + string(vrfctTokenBytes) + "&id=" + fmt.Sprintf("%d", vrfctTokenId) 
        //send email verification link
        err = utils.SendEmailVrfctEmail(displayName, email, vrfctLinkStr)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }else {
        http.Error(w, "error parsing claims from refresh token", http.StatusInternalServerError)
        return
    }
}
