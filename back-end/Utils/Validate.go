package Utils

import(
	"fmt"
	"strings"
)

func ValidateTask(Task string) error {
	if len(Task) < 5 {
		return fmt.Errorf("Task cannot be less than 5 characters")
	}

	if strings.HasPrefix(Task, " ") {
		return fmt.Errorf("Task cannot start with a space")
	}
	return nil
}
