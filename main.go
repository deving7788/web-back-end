package main

import (
    "net/http"
    "log"
    "fmt"
    _ "github.com/lib/pq"
    "web-back-end/database"
    "web-back-end/handlers"
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
    FileHandler := http.FileServer(http.Dir(staticDir))
    muxStatic := http.NewServeMux()
    muxStatic.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        FileHandler.ServeHTTP(w, r)
    })
    go func() {
        fmt.Println("static server listening on 192.168.50.169:8010")
        log.Fatal(http.ListenAndServe("192.168.50.169:8010", muxStatic))
    }()

    //create and run api server
    mux := http.NewServeMux()
    mux.HandleFunc("/api/user/signup", handlers.UserSignupHandler)
    mux.HandleFunc("/api/user/email-vrfct", handlers.SendEmailVrfctHandler)
    mux.HandleFunc("/api/user/email-cfmt", handlers.EmailCfmtHandler)
    mux.HandleFunc("/api/user/login", handlers.UserLoginHandler)
    mux.HandleFunc("/api/user/panel", handlers.AuthenticationHandler)
    mux.HandleFunc("/api/user/panel/change-display-name", handlers.ChangeDisplayNameHandler)
    mux.HandleFunc("/api/user/panel/change-email", handlers.ChangeEmailHandler)
    mux.HandleFunc("/api/user/panel/change-password", handlers.ChangePasswordHandler)
    mux.HandleFunc("/api/user/panel/delete-account", handlers.DeleteAccountHandler)

    fmt.Println("api server listening on localhost:8080")
    log.Fatal(http.ListenAndServe("127.0.0.1:8080", mux))
}

