package database
import (
  "database/sql"
  _ "github.com/lib/pq"
)

var Blogdb *sql.DB

func ConnectDB() (*sql.DB, error) {

  connStr := "user=postgres password=1234 dbname=blogdb sslmode=require"
  db, err := sql.Open("postgres", connStr)
  //err to be handled after invocation of ConnectDB
  if err != nil {
    return nil, err
  }

  return db, nil
}
