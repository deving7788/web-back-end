package midware

import (
    "net/http"
    "fmt"
    "web-back-end/auth"
    "web-back-end/custypes"
)

func SetAccessCookie(w http.ResponseWriter, userToken custypes.UserToken) (error) {

    accessToken, err := auth.CreateAccessToken(userToken)
    if err != nil {
        return fmt.Errorf("error creating access token in SetAccessCookie %v\n", err)
    }

    duration:= 20 * 60 * 1
    jwtCookie := http.Cookie {
        Name: "accessToken",
        Value: accessToken,
        Path: "/",
        MaxAge: duration,
        SameSite: http.SameSiteNoneMode,
        HttpOnly: true,
        Secure: true,
    }

    http.SetCookie(w, &jwtCookie)
    return nil
}
