package database

import (
    "database/sql"
)

func GetRoleById(id int, db *sql.DB) (string, error) {
    queryStr := "SELECT role FROM blog.users WHERE user_id = $1"
    var role string
    err := db.QueryRow(queryStr, id).Scan(&role)
    switch {
    case err == sql.ErrNoRows:
        return "norow", err
    case err != nil:
        return "", err
    default:
        return role, nil
    }

}
