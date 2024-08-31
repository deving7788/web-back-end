package database

import (
    "database/sql"
    "time"
)

func StoreEmailVrfctToken(vrfctTokenBytes []byte, userId int, db *sql.DB) (int, error) {
    var vrfctToken string
    var tokenId int
    expiry_time := time.Now().Add(time.Minute * 15).Format(time.DateTime) 

    dbStr := "SELECT vrfct_token FROM blog.vrfct_tokens WHERE user_id = $1"
    err := db.QueryRow(dbStr, userId).Scan(&vrfctToken)
    if err != nil {
        switch {
        case err == sql.ErrNoRows: 
            dbStr = "INSERT INTO blog.vrfct_tokens (vrfct_token, user_id, expiry_time) VALUES ($1, $2, $3) RETURNING vrfct_token_id"
            err = db.QueryRow(dbStr, string(vrfctTokenBytes), userId, expiry_time).Scan(&tokenId)
        default:
            return -1, err
        }
    }else {
        dbStr = "UPDATE blog.vrfct_tokens SET (vrfct_token, expiry_time) = ($1, $2) WHERE user_id = $3 RETURNING vrfct_token_id"
        err = db.QueryRow(dbStr, string(vrfctTokenBytes), expiry_time, userId).Scan(&tokenId)
    }

    if err != nil {
        return -1, err
    }else {
        return tokenId, nil
    }
}
