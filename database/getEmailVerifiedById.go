package database

import (
    "database/sql"
)

func GetEmailVerifiedById(id int, db *sql.DB) (bool, error) {
    var emailVerified bool
    queryStr := "SELECT email_verified FROM blog.users WHERE user_id = $1"
    err := db.QueryRow(queryStr, id).Scan(&emailVerified)
    if err != nil {
        return false, err
    }else {
        return emailVerified, nil
    }
}
