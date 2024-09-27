package handlers

import (
    "io"
    "errors"
    "encoding/json"
    "net/http"
    "strings"
    "strconv"
    "os"
    "web-back-end/midware"
    "web-back-end/database"
    "web-back-end/auth"
    "web-back-end/utils"
    "github.com/golang-jwt/jwt/v5"
)

func ChangePasswordHandler(w http.ResponseWriter, r *http.Request) {
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
    //unmarshal received passwords
    type BodyPasswords struct {
        OldPassword string
        NewPassword string
    }
    var passwords BodyPasswords
    err = json.Unmarshal(bodyBytes, &passwords)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

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
                //retrieve password from database
                storedPassword, err := database.GetPasswordById(userId, database.Blogdb) 
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError) 
                    return
                }
                //verify password and send bad request if failed
                err = utils.VerifyPassword(storedPassword, passwords.OldPassword)
                if err != nil {
                    http.Error(w, "wrong password", http.StatusBadRequest)
                    return
                }
                //hash password
                costStr := os.Getenv("COST")
                cost, err := strconv.Atoi(costStr)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                hashedPasswordStr, err := utils.HashPassword(passwords.NewPassword, cost)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                //update password in database
                err = database.UpdatePasswordById(userId, hashedPasswordStr, database.Blogdb)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                return
            }else {
                http.Error(w, "error parsing user id from access token", http.StatusInternalServerError)
                return
            }
        }else {
            http.Error(w, "error parsing claims from access token", http.StatusInternalServerError)
            return
        }
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
            //retrieve password from database
            storedPassword, err := database.GetPasswordById(userId, database.Blogdb) 
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError) 
                return
            }
            //verify password and send bad request if failed
            err = utils.VerifyPassword(storedPassword, passwords.OldPassword)
            if err != nil {
                http.Error(w, "wrong password", http.StatusBadRequest)
                return
            }
            //hash password
            costStr := os.Getenv("COST")
            cost, err := strconv.Atoi(costStr)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            hashedPasswordStr, err := utils.HashPassword(passwords.NewPassword, cost)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            //update password in database
            err = database.UpdatePasswordById(userId, hashedPasswordStr, database.Blogdb)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }
            return
        }else {
            http.Error(w, "error parsing user id from refresh token", http.StatusInternalServerError)
            return
        }
    }else {
        http.Error(w, "error parsing claims from refresh token", http.StatusInternalServerError)
        return
    }
}
