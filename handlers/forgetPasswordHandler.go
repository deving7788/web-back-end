package handlers

import (
    "fmt"
    "net/http"
    "os"
    "web-back-end/database"
    "web-back-end/utils"
)

func ForgetPasswordHandler(w http.ResponseWriter, r *http.Request) {
    //get email address from request url
    email := r.URL.Query()["email"][0]
    //get user id from database
    userId, err := database.GetIdByEmail(email, database.Blogdb)
    if err != nil {
        return
    }
    //generate an array of 30 random bytes
    prTokenBytes := utils.GenerateRandomBytes(30) 
    prTokenId, err := database.StorePrToken(prTokenBytes, userId, database.Blogdb)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //generate a password reset link
    apiAddress := os.Getenv("API_HOST")
    baseStr := fmt.Sprintf("%s/api/user/pr-page?", apiAddress)

    prLinkStr := baseStr + "token=" + string(prTokenBytes) + "&id=" + fmt.Sprintf("%d", prTokenId )
    //send password reset email
    err = utils.SendPasswordResetEmail(email, prLinkStr)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
