package main

import (
    "net/http"
    "log"
    "os"
    "fmt"
    "time"
    _ "github.com/lib/pq"
    "web-back-end/database"
    "web-back-end/handlers"
    "web-back-end/midware"
    "web-back-end/rateLimiter"
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

    rl := rateLimiter.NewRateLimiter(100, 60*time.Second, 120*time.Second)

    //create and run api server
    mux := http.NewServeMux()
    mux.HandleFunc("/api/user/signup", handlers.UserSignupHandler)
    mux.HandleFunc("/api/user/email-vrfct", handlers.SendEmailVrfctHandler)
    mux.HandleFunc("/api/user/email-cfmt", handlers.EmailCfmtHandler)
    mux.HandleFunc("/api/user/login", handlers.UserLoginHandler)
    mux.HandleFunc("/api/user/forget-password", handlers.ForgetPasswordHandler)
    mux.HandleFunc("/api/user/pr-page", handlers.SendPasswordResetPageHandler)
    mux.HandleFunc("/api/user/handle-pr", handlers.ResetPasswordHandler)
    mux.HandleFunc("/api/user/panel", handlers.AuthenticationHandler)
    mux.HandleFunc("/api/user/panel/change-display-name", handlers.ChangeDisplayNameHandler)
    mux.HandleFunc("/api/user/panel/change-email", handlers.ChangeEmailHandler)
    mux.HandleFunc("/api/user/panel/change-password", handlers.ChangePasswordHandler)
    mux.HandleFunc("/api/user/panel/delete-account", handlers.DeleteAccountHandler)
    mux.HandleFunc("/api/blog/featured-articles", handlers.GetFeaturedArticlesHandler)
    mux.HandleFunc("/api/blog/article-titles", handlers.GetArticleTitlesHandler)
    mux.HandleFunc("/api/blog/article", handlers.GetArticleHandler)

    handler := midware.RateLimit(mux, rl)
    handler = midware.HandlePreflight(handler)
    handler = midware.SetCors(handler)

    goServerPort := os.Getenv("GO_SERVER_PORT")
    fmt.Println("go api server listening on: ", goServerPort)
    log.Fatal(http.ListenAndServe(goServerPort, handler))
}

