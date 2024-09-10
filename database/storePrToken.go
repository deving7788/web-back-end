package database

import (
    "database/sql"
    "time"
)

func StorePrToken(prTokenBytes []byte, userId int, db *sql.DB) (int, error) {
    var prToken string
    var tokenId int
    expiry_time := time.Now().Add(time.Minute * 15).Format(time.DateTime)

    dbStr := "SELECT pr_token FROM blog.pr_tokens WHERE user_id = $1"
    err := db.QueryRow(dbStr, userId).Scan(&prToken)
    if err != nil {
        switch {
        case err == sql.ErrNoRows:
            dbStr = "INSERT INTO blog.pr_tokens (pr_token, user_id, expiry_time) VALUES ($1, $2, $3) RETURNING pr_token_id"
            err = db.QueryRow(dbStr, string(prTokenBytes), userId, expiry_time).Scan(&tokenId)
        default:
            return -1, err
        }
    }else {
        dbStr = "UPDATE blog.pr_tokens SET (pr_token, expiry_time) = ($1, $2) WHERE user_id = $3 RETURNING pr_token_id"
        err = db.QueryRow(dbStr, string(prTokenBytes), expiry_time, userId).Scan(&tokenId)
    }

    if err != nil {
        return -1, err
    }else {
        return tokenId, nil
    }
}

