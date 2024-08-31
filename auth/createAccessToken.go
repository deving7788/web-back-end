package auth

import (
    "github.com/golang-jwt/jwt/v5"
    "fmt"
    "time"
    "web-back-end/custypes"
    "web-back-end/utils"
)

func CreateAccessToken(userToken custypes.UserToken) (string, error) {
    var (
        token *jwt.Token
        tokenStr string
    )
    
    key, err := utils.ReadEnv("JWT_KEY")
    if err != nil {
        return "", fmt.Errorf("error reading jwtkey in CreateAccessToken %v\n", err)
    }

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

    tokenStr, err = token.SignedString([]byte(key)) 
    if err != nil {
        return "", fmt.Errorf("error signing access jwt token: %v\n", err)
    }
    return tokenStr, nil
}
