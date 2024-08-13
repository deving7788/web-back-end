package database

import (
  "database/sql"
  "fmt"
)


func CheckAccountNameTaken(accName string, db *sql.DB) (bool, error) {
  
  var name string
  str := "SELECT account_name FROM blog.users WHERE account_name = $1"
  err := db.QueryRow(str, accName).Scan(&name)
  
  switch {
    case err == sql.ErrNoRows:
      return false, nil
    case err != nil:
      return true, fmt.Errorf("error checking if account name taken %v\n", err)
    default:
      return true, nil
  }
}
