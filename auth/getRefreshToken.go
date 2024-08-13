package auth

import (
    "net/http"
    "strings"
    "fmt"
)

func GetRefreshToken(r *http.Request) string {
    headers := r.Header
    fmt.Println("headers are: ", headers)
    authHeaders := headers["Authorization"]
    var refreshTokenStr = ""
    for _, authHeader := range authHeaders {
        if strings.Contains(authHeader, "Bearer ") {
            refreshTokenStr = strings.Split(authHeader, "Bearer ")[1]
        }
    }
    fmt.Println("refresh token: ", refreshTokenStr)
    return refreshTokenStr
}
