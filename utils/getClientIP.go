package utils

import (
    "net/http"
    "net"
    "strings"
)

func GetClientIP(r *http.Request) string {
    // Check for X-Forwarded-For
    forwarded := r.Header.Get("X-Forwarded-For")
    if forwarded != "" {
        // Extract the first IP
        return strings.Split(forwarded, ",")[0]
    }

    // Fall back to direct IP from RemoteAddr if no header is found
    ip, _, _ := net.SplitHostPort(r.RemoteAddr)
    return ip
}
