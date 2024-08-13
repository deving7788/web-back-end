package database
import (
    "database/sql"
)

func GetPasswordById(id int, db *sql.DB) (string, error) {

    var password string
    queryStr := "SELECT password FROM blog.users WHERE user_id = $1"
    err := db.QueryRow(queryStr, id).Scan(&password)
    if err != nil {
        return "", err
    }else {
        return password, nil
    }
}
