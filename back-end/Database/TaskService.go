package Database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"todoList/Utils"
)
//Find Tasks
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
//List Tasks
func ListAllTasks(UserToken string) ([]Tasks, error) {
	query := "SELECT * FROM tasks WHERE UserToken = $1"
    rows, err := db.Query(query, UserToken)
    if err != nil {
        return nil, err
    }

    defer rows.Close()

    tasks := []Tasks{}
    for rows.Next() {
        var task Tasks
        err := rows.Scan(&task.Token, &task.Task, &task.Status, &task.UserToken)
        if err != nil {
            return nil, err
        }
        tasks = append(tasks, task)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }
    return tasks, nil
}
//Create Task
func CrateTask(UserToken, Task string) error {
	Token := Utils.GenerateToken(31)
	query := "INSERT INTO tasks (Token, Task, Status, UserToken) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(query,Token,Task,false,UserToken)
	if err != nil {
		return err
	}
	return nil
}
//Delete task
func DeleteTask(Token string) error {
	query := "DELETE FROM tasks WHERE Token=$1"
	_, err := db.Exec(query,Token)
	if err != nil {
		return err
	}
	return nil
}
//Delete All Tasks

func DeleteAllTasks(UserToken string) error {
	query := "DELETE FROM tasks WHERE UserToken=$1"
	_, err := db.Exec(query,UserToken)
	if err != nil {
		return err
	}
	return nil
}
//Update Task
func UpdateTaskStatus(Token string) error {
	Task , err := FindTask(Token)
	if err != nil {
		return err
	}

	query := "UPDATE tasks SET Status=$1 WHERE Token=$2"
	switch Task.Status {
		case true:
			_, err = db.Exec(query,false,Token)
			if err != nil {
				return err
			}
		break;
		case false:
			_, err := db.Exec(query,true,Token)
			if err != nil {
				return err
			}
		break;
	}
	return nil
}
