package database

import (
    "database/sql"
    "fmt"
)

func MarkEmailVerified(tokenId int, db *sql.DB) error {
    var userId int
    queryStr := "SELECT user_id FROM blog.vrfct_tokens WHERE vrfct_token_id = $1"
    err := db.QueryRow(queryStr, tokenId).Scan(&userId) 
    if err != nil {
        return err
    }

    queryStr = "UPDATE blog.users SET email_verified = true WHERE user_id = $1"
    result, err := db.Exec(queryStr, userId)
    if err != nil {
        fmt.Println("%v", err)
    }

    rows, err := result.RowsAffected()
    if err != nil {
        return err
    }

    if rows != 1 {
        fmt.Println("more than 1 row was affected in MarkEmailVerified function")
    }

    return nil
}
