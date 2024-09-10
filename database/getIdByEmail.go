package database
import (
    "database/sql"
)

func GetIdByEmail(email string, db *sql.DB) (int, error) {
    var id int
    queryStr := "SELECT user_id FROM blog.users WHERE email= $1"
    err := db.QueryRow(queryStr, email).Scan(&id)
    if err != nil {
        return -1, err
    }else {
        return id, nil
    }

}
