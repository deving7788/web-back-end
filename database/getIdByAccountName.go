package database
import (
    "database/sql"
)

func GetIdByAccountName(accName string, db *sql.DB) (int, error) {
    
    var id int
    err := db.QueryRow("SELECT user_id FROM blog.users WHERE account_name = $1", accName).Scan(&id)
    switch {
        case err == sql.ErrNoRows:
            return -1, err
        case err != nil:
            return -2, err
        default:
            return id, nil 
    }

}
