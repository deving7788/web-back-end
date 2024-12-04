package midware

import (
  "net/http"
  "os"
)

func SetCors(next http.Handler) http.Handler {
    return http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            allowedOrigin := os.Getenv("ALLOWED_ORIGIN")
            w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
            w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, DELETE, PATCH, PUT")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Cookie")
            w.Header().Set("Access-Control-Allow-Credentials", "true")

            path := r.URL.Path

            switch path {
            case "/api/user/email-cfmt":
                w.Header().Set("Content-Type", "text/html; charset=utf-8")
            case "/api/user/forget-password":
                w.Header().Set("Content-Type", "text/html; charset=utf-8")
            case "/api/user/pr-page":
                w.Header().Set("Content-Type", "text/html; charset=utf-8")
            default:
                w.Header().Set("Content-Type", "application/json")
            }

            next.ServeHTTP(w, r)
        })
}
