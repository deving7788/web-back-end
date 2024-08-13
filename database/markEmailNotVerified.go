package database

import (
    "database/sql"
)

func MarkEmailNotVerified(userId int, db *sql.DB) (bool, error) {
    var emailVerified bool
    queryStr := "UPDATE blog.users SET email_verified = false WHERE user_id = $1 RETURNING email_verified"
    err := db.QueryRow(queryStr, userId).Scan(&emailVerified)
    if err != nil {
        return false, err
    }else {
        return emailVerified, nil
    }
}
