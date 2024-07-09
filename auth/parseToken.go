package auth

import (
    "fmt"
    "github.com/golang-jwt/jwt/v5"
)

func ParseToken(tokenStr string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("err Unexpected access token signing method in UserPanelHandler: %v", token.Header["alg"])
        }
        jwtkey, err := ReadJwtKey()
        if err != nil {
            return nil, err 
        }
        return *jwtkey, nil
    })
    if err != nil {
        return nil, err
    }else {
        return token, nil
    }
    
}
