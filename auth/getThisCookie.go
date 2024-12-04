package auth
import (
    "net/http"
)

func GetThisCookie(cookieName string, r *http.Request) (string) {
    cookies := r.Cookies()
    for _, cookie := range cookies {
        if cookieName == cookie.Name {
            thisCookie := cookie.Value
            return thisCookie
        }
    }
    return ""
}
