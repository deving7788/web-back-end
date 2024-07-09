package database
import (
    "database/sql"
    "fmt"
)

func GetPasswordByAccountName(accName string, db *sql.DB) (string, error) {

    var password string
    queryStr := "SELECT password FROM blog.users WHERE account_name = $1"
    err := db.QueryRow(queryStr, accName).Scan(&password)
    switch {
    case err == sql.ErrNoRows:
        return "noUser", nil
    case err != nil:
        return "", fmt.Errorf("query error in GetPasswordByAccName: %v", err)
    default:
        return password, nil
    }
}
