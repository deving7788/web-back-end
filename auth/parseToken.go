package auth

import (
    "fmt"
    "web-back-end/utils"
    "github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenStr string) (*jwt.Token, error) {
    jwtkeyStr, err := utils.ReadEnv("JWT_KEY")
    jwtkey := []byte(jwtkeyStr)
    if err != nil {
        return nil, err 
    }

    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("err Unexpected access token signing method in UserPanelHandler: %v", token.Header["alg"])
        }
        return jwtkey, nil
    })

    if err != nil {
        return nil, err
    }else {
        return token, nil
    }
    
}
