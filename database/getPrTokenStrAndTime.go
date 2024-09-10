package database

import (
    "database/sql"
    "time"
)

func GetPrTokenStrAndTime(prTokenId int, db *sql.DB) (string, time.Time, error) {
    type queryResult struct {
        prToken string
        expiryTime time.Time
    }
    var qr queryResult

    queryStr := "SELECT pr_token, expiry_time FROM blog.pr_tokens WHERE pr_token_id = $1"
    err := db.QueryRow(queryStr, prTokenId).Scan(&qr.prToken, &qr.expiryTime)
    
    if err != nil {
        return "", time.Time{}, err
    }else {
        return qr.prToken, qr.expiryTime, nil
    }
}
