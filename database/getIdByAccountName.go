package database
import (
    "database/sql"
)

func GetIdByAccountName(accName string, db *sql.DB) (int, error) {
    
    var id int
    err := db.QueryRow("SELECT user_id FROM blog.users WHERE account_name = $1", accName).Scan(&id)
    if err != nil {
        return -1, err
    }else {
        return id, nil
    }
}
