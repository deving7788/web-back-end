package custypes

import (
    "time"
)

type User struct {
    UserId int `json:"userId,omitempty"`
    AccountName string `json:"accountName,omitempty"`
    DisplayName string `json:"displayName,omitempty"`
    Role string `json:"role,omitempty"`
    Email string `json:"email,omitempty"`
    Password string `json:"password,omitempty"`
    CreatedAt time.Time `json:"createdAt,omitempty"`
    EmailVerified bool `json:"emailVerified,omitempty"`
}

type UserLogin struct {
    AccountName string `json:"accountName,omitempty"`
    Password string `json:"password,omitempty"`
}

type UserToken struct {
    UserId int `json:"userId,omitempty"`
    AccountName string `json:"accountName,omitempty"`
    DisplayName string `json:"displayName,omitempty"`
    Role string `json:"role,omitempty"`
    Email string `json:"email,omitempty"`
    EmailVerified bool `json:"emailVerified,omitempty"`
}
