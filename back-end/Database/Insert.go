package Database

import(
	"todoList/Utils"
)

func CrateTask(UserToken, Task string) error {
	Token := Utils.GenerateToken(31)
	query := "INSERT INTO tasks (Token, Task, Status, UserToken) VALUES ($1, $2, $3, $4)"
	_, err := db.Exec(query,Token,Task,false,UserToken)
	if err != nil {
		return err
	}
	return nil
}