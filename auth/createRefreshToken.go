package auth

import (
    "github.com/golang-jwt/jwt/v5"
    "web-back-end/custypes"
    "fmt"
    "time"
)

func CreateRefreshToken(userToken custypes.UserToken) (string, error) {
    var token *jwt.Token
    var tokenStr string
    
    key, err := ReadJwtKey()
    if err != nil {
        return "", fmt.Errorf("error reading jwtkey in CreateRefreshToken %v\n", err)
    }
    
    createdDate := time.Now()
    numericCreatedDate := jwt.NewNumericDate(createdDate)
    expiryDate := time.Now().Add(time.Hour * 24)
    numericExpiryDate := jwt.NewNumericDate(expiryDate)
    token = jwt.NewWithClaims(jwt.SigningMethodHS256,
        jwt.MapClaims {
        "iss": "web-blog",
        "sub": userToken.AccountName,
        "iat": numericCreatedDate,
        "exp": numericExpiryDate, 
        "userId": userToken.UserId,
        "displayName": userToken.DisplayName,
        "role": userToken.Role,
        "email": userToken.Email,
        "emailVerified": userToken.EmailVerified,
        })

    tokenStr, err = token.SignedString(key) 
    if err != nil {
        return "", fmt.Errorf("error signing refresh jwt token: %v\n", err)
    }
    return tokenStr, nil

}
