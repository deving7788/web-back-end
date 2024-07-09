package database

import (
    "database/sql"
)

func ChangeEmailById(userId int, newEmail string, db *sql.DB) (string, error) {
    var email string
    queryStr := "UPDATE blog.users SET email = $1 WHERE user_id = $2 RETURNING email"
    err := db.QueryRow(queryStr, newEmail, userId).Scan(&email) 
    switch {
    case err != nil:
        return "", err
    default:
        return email, nil
    }
}
