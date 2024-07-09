package midware

import (
  "net/http"
)

func SetCors(w http.ResponseWriter) {
  w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5000")
  w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, DELETE, PATCH, PUT")
  w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Cookie")
  w.Header().Set("Access-Control-Allow-Credentials", "true")
}
