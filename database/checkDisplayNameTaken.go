package database

import (
  "database/sql"
  "fmt"
)

func CheckDisplayNameTaken(disName string, db *sql.DB) (bool, error) {
  
  var name string
  str := "SELECT display_name FROM blog.users WHERE display_name = $1"
  err := db.QueryRow(str, disName).Scan(&name)
  
  switch {
    case err == sql.ErrNoRows:
      return false, nil
    case err != nil:
      return true, fmt.Errorf("error checking if display name taken %v\n", err)
    default:
      return true, nil
  }
}
