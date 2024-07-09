package database 

import (
    "fmt"
    "database/sql"
    "time"
    "web-back-end/custypes"
)

func RegisterNewUser(user *custypes.User, db *sql.DB,) (int, error) {
    var id int 
    createdTime := time.Now().Format(time.DateTime)
    dbStr := "INSERT INTO blog.users (account_name, display_name, role, email, password, created_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING user_id"
    err := db.QueryRow(dbStr, user.AccountName, user.DisplayName, user.Role, user.Email, user.Password, createdTime).Scan(&id)
    if err != nil {
        return -1, fmt.Errorf("error registering new user: %v\n", err)
    }

    return id, nil
}
