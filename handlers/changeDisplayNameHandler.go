package handlers

import (
    "io"
    "errors"
    "encoding/json"
    "net/http"
    "strings"
    "web-back-end/midware"
    "web-back-end/custypes"
    "web-back-end/database"
    "web-back-end/auth"
    "github.com/golang-jwt/jwt/v5"
)

func ChangeDisplayNameHandler(w http.ResponseWriter, r *http.Request) {
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
    
    //unmarshal received displayName
    type NewDisplayName struct {
        DisplayName string
    }
    var newDisplayName NewDisplayName
    err = json.Unmarshal(bodyBytes, &newDisplayName)
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
                userId := int (userIdFloat)
                //change displayName in database
                displayName, err := database.ChangeDisplayNameById(userId, newDisplayName.DisplayName, database.Blogdb)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                //assign new display name to response body
                resBody.DisplayName = displayName
            }else {
                http.Error(w, "error parsing user id from claims of access token", http.StatusInternalServerError)
                return
            }
        }else {
            http.Error(w, "error parsing claims from access token", http.StatusInternalServerError)
            return
        }
        //create and send back resoponse body
        resBodyJson, err := json.Marshal(resBody)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        _, err = io.WriteString(w, string(resBodyJson))
        if err !=nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        return
    }

    //get refresh token
    headers := r.Header
    authHeaders := headers["Authorization"]
    var refreshTokenStr string
    for _, authHeader := range authHeaders {
        if strings.Contains(authHeader, "Bearer") {
            refreshTokenStr = strings.Split(authHeader, "Bearer ")[1]
        }
    }
    //parse refresh token
    refreshToken, err := auth.ParseToken(refreshTokenStr)
    if err != nil {
        switch {
        case errors.Is(err, jwt.ErrTokenExpired):
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
            //change displayName in database
            displayName, err := database.ChangeDisplayNameById(userId, newDisplayName.DisplayName, database.Blogdb)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            //create and send back response body
            resBody.DisplayName = displayName
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
    }
}
