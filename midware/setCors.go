package midware

import (
  "net/http"
  "web-back-end/utils"
)

func SetCors(w http.ResponseWriter) {
  allowedOrigin, _ := utils.ReadEnv("ALLOWED_ORIGIN")
  w.Header().Set("Access-Control-Allow-Origin", allowedOrigin)
  w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, DELETE, PATCH, PUT")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Cookie")
  w.Header().Set("Access-Control-Allow-Credentials", "true")
}
