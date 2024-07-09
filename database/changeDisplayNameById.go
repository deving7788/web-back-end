package database

import (
    "database/sql"
)

func ChangeDisplayNameById(userId int, newDisplayName string, db *sql.DB) (string, error) {
    var displayName string
    queryStr := "UPDATE blog.users SET display_name = $1 WHERE user_id = $2 RETURNING display_name"
    err := db.QueryRow(queryStr, newDisplayName, userId).Scan(&displayName)
    switch {
    case err != nil:
        return "", err
    default:
        return displayName, nil
    }
}
