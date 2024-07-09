package handlers
import (
    "net/http"
    "errors"
    "io"
    "encoding/json"
    "strings"
    "web-back-end/midware"
    "web-back-end/auth"
    "web-back-end/custypes"
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
    //get accessToken cookie and parse access Token
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
                http.Error(w, "error parse accountName from claims of access token", http.StatusInternalServerError)
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
                http.Error(w, "err parse user id from claims of access token", http.StatusInternalServerError)
                return
            }
            //parse display name
            displayName, ok := claims["displayName"].(string)
            if ok {
                userToken.DisplayName = displayName
                resBody.DisplayName = displayName
            }else {
                http.Error(w, "error parse displayName from claims of access token", http.StatusInternalServerError)
                return
            }
            //parse role
            role, ok := claims["role"].(string)
            if ok {
                userToken.Role = role
                resBody.Role = role
            }else {
                http.Error(w, "error parse role from claims of access token", http.StatusInternalServerError)
                return
            }
            //parse email
            email, ok := claims["email"].(string)
            if ok {
                userToken.Email = email
                resBody.Email = email
            }else {
                http.Error(w, "error parsing email from claims of access token", http.StatusInternalServerError)
                return
            }
        }else {
            http.Error(w, "error parse claims of access token", http.StatusInternalServerError)
            return
        }
        //write access token into cookie header
        err = midware.SetAccessCookie(w, &userToken) 
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
        _, err = io.WriteString(w, string(resBodyJson));
        if err != nil {
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
        if strings.Contains(authHeader, "Bearer ") {
            refreshTokenStr = strings.Split(authHeader, "Bearer ")[1]
        }
    }

    //parse refresh token
    refreshToken, err := auth.ParseToken(refreshTokenStr)

    //handle parsing error and expired token
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
        //parse display name
        displayName, ok := claims["displayName"].(string) 
        if ok {
            userToken.DisplayName = displayName
            resBody.DisplayName = displayName
        }else {
            http.Error(w, "error parsing displayName claim of refresh token", http.StatusInternalServerError)
            return
        }
        //parse role
        role, ok := claims["role"].(string) 
        if ok {
            userToken.Role = role
            resBody.Role = role
        }else {
            http.Error(w, "error parsing role claim of refresh token", http.StatusInternalServerError)
            return
        }
        //parse email
        email, ok := claims["email"].(string) 
        if ok {
            userToken.Email = email
            resBody.Email = email
        }else {
            http.Error(w, "error parsing email claim of refresh token", http.StatusInternalServerError)
            return
        }
    }else {
        http.Error(w, "error parsing claims from refresh token", http.StatusInternalServerError)
    }

    //write access token into cookie header
    err = midware.SetAccessCookie(w, &userToken) 
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    //create refresh token
    refreshTokenStr, err = auth.CreateRefreshToken(&userToken)
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
