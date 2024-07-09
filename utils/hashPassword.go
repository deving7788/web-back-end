package utils

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string, cost int) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*password), cost)
    if err != nil {
        return "", fmt.Errorf("err hashing password: %v", err)
    }
    
    return string(hashedPassword), nil
}
