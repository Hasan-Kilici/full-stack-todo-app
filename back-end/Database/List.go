package Database

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