package database

import (
  "fmt"
  "database/sql"
)

func DeleteUserByAccountName(accName string, db *sql.DB) (error) {

  _, err := db.Exec("DELETE FROM blog.users WHERE account_name = $1", accName)
  if err != nil {
    return fmt.Errorf("error deleting user %v: %v\n", accName, err) 
  }else {
    return nil
  }
}
