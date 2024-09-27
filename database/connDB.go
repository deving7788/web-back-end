package database
import (
    "os"
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)

var Blogdb *sql.DB

func ConnectDB() (*sql.DB, error) {

    dbPassword := os.Getenv("DB_PASSWORD")
    connStr := fmt.Sprintf("user=postgres password=%s dbname=blogdb sslmode=require", dbPassword)
    db, err := sql.Open("postgres", connStr)
    //err to be handled after invocation of ConnectDB
    if err != nil {
        return nil, err
    }

    return db, nil
}
