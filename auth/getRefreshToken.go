package auth

import (
    "net/http"
    "strings"
)

func GetRefreshToken(r *http.Request) string {
    headers := r.Header
    authHeaders := headers["Authorization"]
    var refreshTokenStr = ""
    for _, authHeader := range authHeaders {
        if strings.Contains(authHeader, "Bearer ") {
            refreshTokenStr = strings.Split(authHeader, "Bearer ")[1]
        }
    }
    return refreshTokenStr
}
