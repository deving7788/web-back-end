package handlers

import (
    "io"
    "encoding/json"
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
    //get login user's password from database
    retrievedPassword, err := database.GetPasswordByAccountName(user.AccountName, database.Blogdb)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    //verify password. Return 401 if password is incorrect
    err = utils.VerifyPassword(&retrievedPassword, &user.Password)
    if err != nil {
        http.Error(w, "password does not match", http.StatusUnauthorized)
        return
    }

    //get user's role, id, email, displayName,
    userId, err := database.GetIdByAccountName(user.AccountName, database.Blogdb) 
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
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

    //create userToken for access and refresh tokens
    var userToken custypes.UserToken
    userToken.UserId = userId
    userToken.AccountName = user.AccountName
    userToken.DisplayName = displayName
    userToken.Role = role
    userToken.Email = email

    //set access token
    err = midware.SetAccessCookie(w, &userToken)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    //create and return refresh token
    refreshToken, err := auth.CreateRefreshToken(&userToken)
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
