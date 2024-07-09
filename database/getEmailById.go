package database

import (
    "database/sql"
)

func GetEmailById(id int, db *sql.DB) (string, error) {
    queryStr := "SELECT email FROM blog.users WHERE user_id = $1"
    var email string
    err := db.QueryRow(queryStr, id).Scan(&email)
    
    switch {
    case err == sql.ErrNoRows: 
        return "norow", err
    case err != nil:
        return "", err
    default:
        return email, nil
    }
}
