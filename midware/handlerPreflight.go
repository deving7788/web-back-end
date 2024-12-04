package midware

import (
    "strings"
    "net/http"
)

func HandlePreflight(next http.Handler) http.Handler {
    return http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            if strings.ToLower(r.Method) == "options" {
                return
            }

            next.ServeHTTP(w, r)
        })
}
