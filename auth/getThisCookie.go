package auth
import (
    "net/http"
    "strings"
)

func GetThisCookie(cookieName string, r *http.Request) (string) {
    headers := r.Header
    cookies := headers["Cookie"]
    thisCookie := ""
    for _, cookie := range cookies {
        if strings.Contains(cookie, cookieName) {
            thisCookie = cookie
        }
    }
    return thisCookie
}
