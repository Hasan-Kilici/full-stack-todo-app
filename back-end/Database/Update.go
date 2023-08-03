package Database

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