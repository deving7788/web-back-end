package handlers

import (
    "net/http"
    "strings"
    "os"
    "io"
    "strconv"
    "encoding/json"
    "web-back-end/database"
    "web-back-end/custypes"
    "web-back-end/midware"
    "web-back-end/auth"
    "web-back-end/utils"
)

func UserSignupHandler(w http.ResponseWriter, r *http.Request) {
    midware.SetCors(w)
    w.Header().Set("Content-Type", "application/json")

    //handle preflight request
    if strings.ToLower(r.Method) == "options" {
        return
    }

    //declear and initialize response body
    var resBody custypes.ResponseBodySignup

    //get request body of json 
    bodyStream := r.Body
    bodyBytes, err := io.ReadAll(bodyStream)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer bodyStream.Close()

    //convert json into struct
    var newUser custypes.User
    err = json.Unmarshal(bodyBytes, &newUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError) 
        return
    }
    
    //hash the password
    costStr := os.Getenv("COST")
    cost, err := strconv.Atoi(costStr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    hashedPasswordStr, err := utils.HashPassword(newUser.Password, cost)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    //replace password with hashed password, and set EmailVerified
    newUser.Password = hashedPasswordStr
    newUser.EmailVerified = false

    //check if account name exists in db
    accTaken, err := database.CheckAccountNameTaken(newUser.AccountName, database.Blogdb) 
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError) 
        return
    }
    if accTaken == true {
        resBody.AccountNameProm = "ACCOUNT_NAME_IS_TAKEN"
    }

    //check if display name exists in db
    disTaken, err := database.CheckDisplayNameTaken(newUser.DisplayName, database.Blogdb)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if disTaken == true {
        resBody.DisplayNameProm = "DISPLAY_NAME_IS_TAKEN"
    }

    //check if email exists in db
    used, err := database.CheckEmailUsed(newUser.Email, database.Blogdb)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    if used == true {
        resBody.EmailProm = "EMAIL_IS_USED"
    }
    //return bad request status on failure of previous three checks
    if accTaken == true || disTaken == true || used == true {
        resBodyJson, err := json.Marshal(resBody)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError) 
            return
        }

        w.WriteHeader(http.StatusBadRequest)
        _, err = io.WriteString(w, string(resBodyJson))
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError) 
            return
        }
        return
    }

    //write new user to database
    userId, err := database.RegisterNewUser(&newUser, database.Blogdb) 
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError) 
        return
    }
    
    //declare and initialize userToken struct for create tokens
    var userToken custypes.UserToken
    userToken.UserId = userId
    userToken.AccountName = newUser.AccountName
    userToken.DisplayName = newUser.DisplayName
    userToken.Role = newUser.Role
    userToken.Email = newUser.Email
    userToken.EmailVerified = false

    //set access token cookie
    err = midware.SetAccessCookie(w, userToken)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotImplemented) 
        return
    }
    
    //create and add refresh token to response body
    refreshTokenStr, err := auth.CreateRefreshToken(userToken)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotImplemented) 
        return
    }

    //prepare response body 
    resBody.AccountNameProm = "OK"
    resBody.DisplayNameProm = "OK"
    resBody.EmailProm = "OK"
    resBody.AccountName = newUser.AccountName
    resBody.DisplayName = newUser.DisplayName
    resBody.Role = newUser.Role
    resBody.Email = newUser.Email
    resBody.RefreshToken = refreshTokenStr
    resBody.EmailVerified = false
    
    //json response body
    resBodyJson, err := json.Marshal(resBody) 
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotImplemented) 
        return
    }

    //return response body
    _, err = io.WriteString(w, string(resBodyJson));
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotImplemented) 
    }
}
