package auth

import (
  "fmt"
  "os"
)

func ReadJwtKey() ([]byte, error) {
  var key []byte 
  key, err := os.ReadFile("auth/jwtkey")
  if err != nil {
    return nil, fmt.Errorf("error reading jwt key file: %v\n", err) 
  }

  return key, nil
}
