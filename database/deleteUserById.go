package database

import (
    "database/sql"
    "fmt"
)

func DeleteUserById(id int, db *sql.DB) (error) {
    dbStr := "DELETE FROM blog.users WHERE user_id = $1"
    result, err := db.Exec(dbStr, id) 
    if err != nil {
        return err
    }
    
    rows, err := result.RowsAffected()
    if err != nil {
        fmt.Printf("an error in DeleteUserById: %v, user id: %v\n", err, id)
    }
    if rows != 1 {
        fmt.Printf("more than one row affected in DeleteUserById, user id: %v\n", id)
    }

    return nil
}
