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
    //create and run postgreSql database connection
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

    //create and run static server goroutine
    staticDir := "/home/wb/myproj/sample/test-app/dist"
    staticHandler := http.FileServer(http.Dir(staticDir))
    go func() {
        fmt.Println("static server listening on port: 8000")
        log.Fatal(http.ListenAndServe(":8000", staticHandler))
    }()

    //create and run api server
    mux := http.NewServeMux()
    mux.HandleFunc("/api/user/signup", handlers.UserSignupHandler)
    mux.HandleFunc("/api/user/login", handlers.UserLoginHandler)
    mux.HandleFunc("/api/user/panel", handlers.AuthenticationHandler)
    mux.HandleFunc("/api/user/panel/change-display-name", handlers.ChangeDisplayNameHandler)
    mux.HandleFunc("/api/user/panel/change-email", handlers.ChangeEmailHandler)

    fmt.Println("api server listening on port: 8080")
    log.Fatal(http.ListenAndServe(":8080", mux))
}

