package database

import (
    "database/sql"
)

func GetEmailById(id int, db *sql.DB) (string, error) {
    queryStr := "SELECT email FROM blog.users WHERE user_id = $1"
    var email string
    err := db.QueryRow(queryStr, id).Scan(&email)

    if err != nil {
        return "", err
    }else {
        return email, nil
    }
}
