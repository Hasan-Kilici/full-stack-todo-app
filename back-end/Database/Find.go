package Database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

func FindTask(Token string) (Tasks, error){
    query := "SELECT * FROM tasks WHERE token = $1"
    row := db.QueryRow(query, Token)

    var tasks Tasks
    err := row.Scan(&tasks.Token, &tasks.Task, &tasks.Status, &tasks.UserToken)
    if err != nil {
        if err == sql.ErrNoRows {
            return Tasks{}, fmt.Errorf("task not found")
        }
        return Tasks{}, err
    }
    return tasks, nil
}