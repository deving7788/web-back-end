package auth

import (
    "github.com/golang-jwt/jwt/v5"
    "fmt"
    "time"
    "web-back-end/custypes"
)

func CreateAccessToken(userToken *custypes.UserToken) (string, error) {
    var (
        token *jwt.Token
        tokenStr string
    )
    
    key, err := ReadJwtKey()
    if err != nil {
        return "", fmt.Errorf("error reading jwtkey in CreateAccessToken %v\n", err)
    }

    createdDate := time.Now()
    numericCreatedDate := jwt.NewNumericDate(createdDate)
    token = jwt.NewWithClaims(jwt.SigningMethodHS256,
        jwt.MapClaims {
        "iss": "web-blog",
        "sub": userToken.AccountName,
        "iat": &numericCreatedDate,
        "userId": userToken.UserId,
        "role": userToken.Role,
        "displayName": userToken.DisplayName,
        "email": userToken.Email,
        })

    tokenStr, err = token.SignedString(*key) 
    if err != nil {
        return "", fmt.Errorf("error signing access jwt token: %v\n", err)
    }
    return tokenStr, nil
}
