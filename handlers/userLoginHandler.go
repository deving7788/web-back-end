package handlers

import (
    "io"
    "encoding/json"
    "database/sql"
    "net/http"
    "strings"
    "web-back-end/midware"
    "web-back-end/custypes"
    "web-back-end/database"
    "web-back-end/utils"
    "web-back-end/auth"
)

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
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

    //convert json into struct
    var user custypes.UserLogin
    err = json.Unmarshal(bodyBytes, &user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //retrieve login user's id from database
    userId, err := database.GetIdByAccountName(user.AccountName, database.Blogdb) 
    if err != nil {
        switch {
        case err == sql.ErrNoRows:
            http.Error(w, "bad request", http.StatusBadRequest)
            return
        default:
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
    }
    //retrieve password from database
    retrievedPassword, err := database.GetPasswordById(userId, database.Blogdb)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    //verify password. Return 400 if password is incorrect
    err = utils.VerifyPassword(retrievedPassword, user.Password)
    if err != nil {
        http.Error(w, "bad request", http.StatusBadRequest)
        return
    }

    //retrieve user's displayName, role, email, emailVerified from database
    displayName, err := database.GetDisplayNameById(userId, database.Blogdb)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    role, err := database.GetRoleById(userId, database.Blogdb)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    email, err := database.GetEmailById(userId, database.Blogdb)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    emailVerified, err := database.GetEmailVerifiedById(userId, database.Blogdb)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    //create userToken for access and refresh tokens
    var userToken custypes.UserToken
    userToken.UserId = userId
    userToken.AccountName = user.AccountName
    userToken.DisplayName = displayName
    userToken.Role = role
    userToken.Email = email
    userToken.EmailVerified = emailVerified

    //set access token
    err = midware.SetAccessCookie(w, userToken)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    //create and return refresh token
    refreshToken, err := auth.CreateRefreshToken(userToken)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //create and send back response body
    var resBody custypes.ResponseBodyLogin
    resBody.AccountName = user.AccountName
    resBody.DisplayName = displayName
    resBody.Role = role
    resBody.Email = email
    resBody.RefreshToken = refreshToken
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
}
