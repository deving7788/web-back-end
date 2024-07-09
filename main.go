package main

import (
    "net/http"
    "log"
    "fmt"
    _ "github.com/lib/pq"
    "web-back-end/handlers"
    "web-back-end/database"
)

func main() {
    var errConnDB error 
    database.Blogdb, errConnDB = database.ConnectDB()
    if errConnDB != nil {
        log.Fatal(errConnDB)
    }
    
    errPing := database.Blogdb.Ping()
    if errPing != nil{
        log.Fatal(errPing)
    }
    fmt.Println("PostgreSql connection established")

    defer database.Blogdb.Close()
    
    http.HandleFunc("/api/user/signup", handlers.UserSignupHandler)
    http.HandleFunc("/api/user/login", handlers.UserLoginHandler)
    http.HandleFunc("/api/user/panel", handlers.AuthenticationHandler)
    http.HandleFunc("/api/user/panel/change-display-name", handlers.ChangeDisplayNameHandler)
    http.HandleFunc("/api/user/panel/change-email", handlers.ChangeEmailHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

