package auth

import (
    "github.com/golang-jwt/jwt/v5"
    "fmt"
    "time"
    "os"
    "web-back-end/custypes"
)

func CreateAccessToken(userToken custypes.UserToken) (string, error) {
    var (
        token *jwt.Token
        tokenStr string
    )
    
    key := os.Getenv("JWT_KEY")

    createdDate := time.Now()
    numericCreatedDate := jwt.NewNumericDate(createdDate)
    token = jwt.NewWithClaims(jwt.SigningMethodHS256,
        jwt.MapClaims {
        "iss": "web-blog",
        "sub": userToken.AccountName,
        "iat": numericCreatedDate,
        "userId": userToken.UserId,
        "displayName": userToken.DisplayName,
        "role": userToken.Role,
        "email": userToken.Email,
        "emailVerified": userToken.EmailVerified,
        })

    tokenStr, err := token.SignedString([]byte(key)) 
    if err != nil {
        return "", fmt.Errorf("error signing access jwt token: %v\n", err)
    }
    return tokenStr, nil
}
