package database

import (
  "database/sql"
  "fmt"
)


func CheckEmailUsed(email string, db *sql.DB) (bool, error) {
  
  var emailDb string
  str := "SELECT email FROM blog.users WHERE email = $1"
  err := db.QueryRow(str, email).Scan(&emailDb)
  
  switch {
    case err == sql.ErrNoRows:
      return false, nil
    case err != nil:
      return false, fmt.Errorf("error checking if email used %v\n", err)
    default:
      return true, nil
  }
}
