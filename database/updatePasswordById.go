package database

import (
    "database/sql"
    "fmt"
)

func UpdatePasswordById(userId int, password string, db *sql.DB) error {
    queryStr := "UPDATE blog.users SET password = $1 WHERE user_id = $2"
    result, err := db.Exec(queryStr, password, userId)

    if err != nil {
        return err
    }

    rows, err := result.RowsAffected()
    if err != nil {
        fmt.Println(err)
    }
    if rows != 1 {
        fmt.Println("more than one row affected in UpdatePasswordById")
    }

    return nil
}

