package database
import (
    "web-back-end/utils"
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

var Blogdb *sql.DB

func ConnectDB() (*sql.DB, error) {

    dbPassword, err := utils.ReadEnv("DB_PASSWORD")
    if err != nil {
        fmt.Printf("err in ConnectDB: %v\n", err)
    }
    connStr := fmt.Sprintf("user=postgres password=%s dbname=blogdb sslmode=require", dbPassword)
    db, err := sql.Open("postgres", connStr)
    //err to be handled after invocation of ConnectDB
    if err != nil {
        return nil, err
    }

    return db, nil
}
