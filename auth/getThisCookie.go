package auth
import (
    "net/http"
)

func GetThisCookie(cookieName string, r *http.Request) (string) {
    //headers := r.Header
    //cookies := headers["Cookie"]
    //fmt.Printf("cookies in GetThisCookie: %v\n", cookies)
    //thisCookie := ""
    //for _, cookie := range cookies {
      //  if strings.Contains(cookie, cookieName) {
       //     thisCookie = cookie
        //}
    //}
    cookies := r.Cookies()
    for _, cookie := range cookies {
        if cookieName == cookie.Name {
            thisCookie := cookie.Value
            return thisCookie
        }
    }
    return ""
}
