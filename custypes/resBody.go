package custypes

type ResponseBodySignup struct {
    AccountNameProm string `json:"accountNameProm,omitempty"`
    DisplayNameProm string `json:"displayNameProm,omitempty"`
    EmailProm string `json:"emailProm,omitempty"`
    AccountName string `json:"accountName,omitempty"`
    DisplayName string `json:"displayName,omitempty"`
    Role string `json:"role,omitempty"`
    Email string `json:"email,omitempty"`
    RefreshToken string `json:"refreshToken,omitempty"`
    EmailVerified bool `json:"emailVerified"`
}

type ResponseBodyUser struct {
    DisplayNameProm string `json:"displayNameProm,omitempty"`
    EmailProm string `json:"emailProm,omitempty"`
    AccountName string `json:"accountName,omitempty"`
    DisplayName string `json:"displayName,omitempty"`
    Role string `json:"role,omitempty"`
    Email string `json:"email,omitempty"`
    RefreshToken string `json:"refreshToken,omitempty"`
    EmailVerified bool `json:"emailVerified"`
}

type ResponseBodyLogin struct {
    AccountName string `json:"accountName,omitempty"`
    DisplayName string `json:"displayName,omitempty"`
    Role string `json:"role,omitempty"`
    Email string `json:"email,omitempty"`
    RefreshToken string `json:"refreshToken,omitempty"`
    EmailVerified bool `json:"emailVerified"`
}

