package utils

import (
    "golang.org/x/crypto/bcrypt"
)

func VerifyPassword(hashedPassword string, password string) error {
    hashedPasswordByte := []byte(hashedPassword)
    passwordByte := []byte(password)
    err := bcrypt.CompareHashAndPassword(hashedPasswordByte, passwordByte)
    return err
}
