package database 

import (
    "database/sql"
)

func GetIdByPrTokenId(prTokenId int, db *sql.DB) (int, error) {
    var userId int
    queryStr := "SELECT user_id FROM blog.pr_tokens WHERE pr_token_id = $1"
    err := db.QueryRow(queryStr, prTokenId).Scan(&userId)
    if err != nil {
        return -1, err
    }else {
        return userId, nil
    }
}
