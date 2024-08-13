package handlers
import (
    "net/http"
    "errors"
    "io"
    "encoding/json"
    "strings"
    "database/sql"
    "web-back-end/midware"
    "web-back-end/auth"
    "web-back-end/custypes"
    "web-back-end/database"
    "github.com/golang-jwt/jwt/v5"
)

func AuthenticationHandler(w http.ResponseWriter, r *http.Request ) {
    midware.SetCors(w)
    w.Header().Set("Content-Type", "application/json")

    //handle preflight request
    if strings.ToLower(r.Method) == "options" {
        return
    }
    //declare response body
    var resBody custypes.ResponseBodyUser
    //declare user token
    var userToken custypes.UserToken
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
        //parse claims from accessToken
        claims, ok := accessToken.Claims.(jwt.MapClaims)
        if ok {
            //parse account name
            accountName, err := claims.GetSubject()
            if err != nil {
                http.Error(w, "error parsing accountName from claims of access token", http.StatusInternalServerError)
                return
            }else {
                userToken.AccountName = accountName
                resBody.AccountName = accountName
            }
            //parse user id
            userIdFloat, ok := claims["userId"].(float64)
            if ok {
                userToken.UserId = int(userIdFloat)
            }else {
                http.Error(w, "err parsing user id from claims of access token", http.StatusInternalServerError)
                return
            }
            //get display name from database
            displayName, err := database.GetDisplayNameById(userToken.UserId, database.Blogdb)
            if err == nil {
                userToken.DisplayName = displayName
                resBody.DisplayName = displayName
            }else {
                if err == sql.ErrNoRows {
                    http.Error(w, err.Error(), http.StatusNotFound)
                    return
                }
                http.Error(w, "error getting display name in access token", http.StatusInternalServerError)
                return
            }
            //get role from database
            role, err := database.GetRoleById(userToken.UserId, database.Blogdb)
            if err == nil {
                userToken.Role = role
                resBody.Role = role
            }else {
                http.Error(w, "error getting role in access token", http.StatusInternalServerError)
                return
            }
            //get email from database
            email, err := database.GetEmailById(userToken.UserId, database.Blogdb)
            if err == nil {
                userToken.Email = email
                resBody.Email = email
            }else {
                http.Error(w, "error getting email in access token", http.StatusInternalServerError)
                return
            }
            //get emailVerified from database
            emailVerified, err := database.GetEmailVerifiedById(userToken.UserId, database.Blogdb)
            if err == nil {
                userToken.EmailVerified = emailVerified
                resBody.EmailVerified = emailVerified
            }else {
                http.Error(w, "error getting emailVerified in access token", http.StatusInternalServerError)
                return
            }
        }else {
            http.Error(w, "error parsing claims of access token", http.StatusInternalServerError)
            return
        }
        //write access token into cookie header
        err = midware.SetAccessCookie(w, userToken) 
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        //json response body
        resBodyJson, err := json.Marshal(resBody)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        //send response body and return
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
        //parse account name
        accountName, err := claims.GetSubject()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }else {
            userToken.AccountName = accountName
            resBody.AccountName = accountName
        }
        //parse user id
        userIdFloat, ok := claims["userId"].(float64) 
        if ok {
            userToken.UserId = int(userIdFloat) 
        }else {
            http.Error(w, "error parsing userId claim of refresh token", http.StatusInternalServerError)
            return
        }
        //get display name from database
        displayName, err := database.GetDisplayNameById(userToken.UserId, database.Blogdb)
        if err == nil {
            userToken.DisplayName = displayName
            resBody.DisplayName = displayName
        }else {
            if err == sql.ErrNoRows {
                http.Error(w, err.Error(), http.StatusNotFound)
                return
            }
            http.Error(w, "error getting display name in refresh token", http.StatusInternalServerError)
            return
        }
        //get role from database
        role, err := database.GetRoleById(userToken.UserId, database.Blogdb)
        if err == nil {
            userToken.Role = role
            resBody.Role = role
        }else {
            http.Error(w, "error getting role in refresh token", http.StatusInternalServerError)
            return
        }
        //get email from database
        email, err := database.GetEmailById(userToken.UserId, database.Blogdb)
        if err == nil {
            userToken.Email = email
            resBody.Email = email
        }else {
            http.Error(w, "error getting email in refresh token", http.StatusInternalServerError)
            return
        }
        //get emailVerified from database
        emailVerified, err := database.GetEmailVerifiedById(userToken.UserId, database.Blogdb)
        if err == nil {
            userToken.EmailVerified = emailVerified
            resBody.EmailVerified = emailVerified
        }else {
            http.Error(w, "error getting emailVerified in  refresh token", http.StatusInternalServerError)
            return
        }
    }else {
        http.Error(w, "error parsing claims from refresh token", http.StatusInternalServerError)
        return
    }

    //write access token into cookie header
    err = midware.SetAccessCookie(w, userToken) 
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //create refresh token
    refreshTokenStr, err = auth.CreateRefreshToken(userToken)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //attach refresh token to response body
    resBody.RefreshToken = refreshTokenStr
    //json and write response body
    resBodyJson, err := json.Marshal(resBody)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    _, err = io.WriteString(w, string(resBodyJson));
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
}
