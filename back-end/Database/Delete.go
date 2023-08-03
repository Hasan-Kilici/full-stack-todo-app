package Database

func DeleteTask(Token string) error {
	query := "DELETE FROM tasks WHERE Token=$1"
	_, err := db.Exec(query,Token)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllTasks(UserToken string) error {
	query := "DELETE FROM tasks WHERE UserToken=$1"
	_, err := db.Exec(query,UserToken)
	if err != nil {
		return err
	}
	return nil
}