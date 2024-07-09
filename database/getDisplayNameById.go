package database

import (
    "database/sql"
)

func GetDisplayNameById(id int, db *sql.DB) (string, error) {
    queryStr := "SELECT display_name FROM blog.users WHERE user_id = $1"
    var displayName string
    err := db.QueryRow(queryStr, id).Scan(&displayName)
    
    switch {
    case err == sql.ErrNoRows: 
        return "norow", err
    case err != nil:
        return "", err
    default:
        return displayName, nil
    }
}
