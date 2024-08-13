package database

import (
    "database/sql"
)

func GetEmailVerifiedById(id int, db *sql.DB) (bool, error) {
    var emailVerified bool
    queryStr := "select email_verified from blog.users where user_id = $1"
    err := db.QueryRow(queryStr, id).Scan(&emailVerified)
    if err != nil {
        return false, err
    }else {
        return emailVerified, nil
    }
}
