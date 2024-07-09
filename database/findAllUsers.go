package database
import (
  "database/sql"
  "fmt"
  "web-back-end/custypes"
)

func FindAllUsers(db *sql.DB) ([]custypes.User, error) {

  var users []custypes.User
  rows, err := db.Query("SELECT user_id, account_name, display_name, email, created_at FROM blog.users")

  if err != nil {
    return nil, fmt.Errorf("An error occurred querying all users in FindAllUsers: %v", err)
  }

  for rows.Next() {
    var tempUser custypes.User
    err := rows.Scan(&tempUser.UserId, &tempUser.AccountName, &tempUser.DisplayName, &tempUser.Email, &tempUser.CreatedAt)
    if err != nil {
      return nil, fmt.Errorf("A scanning error in FindAllUsers: %v", err)
    }
    users = append(users, tempUser)
  }
  
  if rows.Err() != nil {
    return nil, fmt.Errorf("A Rows.Err() error in FindAllUsers")
  }

  defer rows.Close()
  
  return users, nil 
}
