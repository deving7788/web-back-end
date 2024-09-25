package main

import (
    "net/http"
    "log"
    "os"
    "fmt"
    _ "github.com/lib/pq"
    "web-back-end/database"
    "web-back-end/handlers"
    "web-back-end/utils"
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
    //staticDir := "/home/wb/myproj/sample/test-app/dist"
    //FileHandler := http.FileServer(http.Dir(staticDir))
    //muxStatic := http.NewServeMux()
    //muxStatic.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
        //FileHandler.ServeHTTP(w, r)
    //})
    //go func() {
        //fmt.Println("static server listening on :8010")
        //log.Fatal(http.ListenAndServe(":8010", muxStatic))
    //}()

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

    goServerPort, err := utils.ReadEnv("GO_SERVER_PORT")
    goServerPort1 := os.Getenv("GO_SERVER_PORT")
    fmt.Printf("port is %v\n", goServerPort)
    fmt.Printf("port1 is %v\n", goServerPort1)
    fmt.Println("go server listening on: ", goServerPort1)

    if err != nil {
        log.Fatal("error reading api address in main: %v\n", err)
    }
    log.Fatal(http.ListenAndServe(goServerPort1, mux))
    //log.Fatal(http.ListenAndServeTLS(goServerPort, "tls/fullchain.pem", "tls/privkey.pem", mux))
}

