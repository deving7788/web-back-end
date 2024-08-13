package database

import (
    "database/sql"
    "time"
)

func GetVrfctTokenStrAndTime(vrfctId int, db *sql.DB) (string, time.Time, error) {
    type queryResult struct {
        vrfctToken string
        expiryTime time.Time
    }
    var qr queryResult

    queryStr := "SELECT vrfct_token, expiry_time FROM blog.vrfct_tokens WHERE vrfct_token_id = $1"
    err := db.QueryRow(queryStr, vrfctId).Scan(&qr.vrfctToken, &qr.expiryTime)

    if err != nil {
        return "", time.Time{}, err
    }else {
        return qr.vrfctToken, qr.expiryTime, nil
    }

}
