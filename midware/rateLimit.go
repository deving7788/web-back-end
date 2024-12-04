package midware


import (
    "net/http"
    "web-back-end/utils"
    "web-back-end/rateLimiter"
)


func RateLimit(next http.Handler, rl *rateLimiter.RateLimiter) http.Handler {
    return http.HandlerFunc(
        func(w http.ResponseWriter, r *http.Request) {
            ip := utils.GetClientIP(r)
            if !rl.IsAllowed(ip) {
                http.Error(w, "Taking a break can be refreshing.", http.StatusTooManyRequests)
                return
            }

            next.ServeHTTP(w, r)
        })
}
