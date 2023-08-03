package Handlers

import (
	"github.com/gofiber/fiber/v2"
	"todoList/Database"
)

func HelloHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"hello": "world",
	})
}

func CreateTask(c *fiber.Ctx) error {
	User := c.Cookies("Token")
	form := new(TaskForm)
	if err := c.BodyParser(form); err != nil {
		return c.Status(400).SendString("Bad request")
	}

	err := Database.CrateTask(User, form.Task)
	if err != nil {
		return c.Status(500).SendString("Task didn't create")
	}

	c.JSON(fiber.Map{
		"message":"task successfuly created",
	})

	return nil
}

func DeleteTask(c *fiber.Ctx) error {
	Token := c.Params("Token")
	err := Database.DeleteTask(Token)
	if err != nil {
		return c.Status(500).SendString("Task didn't delete")
	}

	c.JSON(fiber.Map{
		"message":"task successfuly deleted",
	})

	return nil
}

func DeleteAllTasks(c *fiber.Ctx) error {
	User := c.Cookies("Token")
	err := Database.DeleteAllTasks(User)
	if err != nil {
		return c.Status(500).SendString("All Tasks didn't delete")
	}

	c.JSON(fiber.Map{
		"message":"all tasks successfuly deleted",
	})

	return nil
}

func UpdateTaskStatus(c *fiber.Ctx) error {
	Token := c.Params("Token")
	err := Database.UpdateTaskStatus(Token)
	if err != nil {
		return c.Status(500).SendString("Task status didn't update")
	}

	c.JSON(fiber.Map{
		"message":"task status successfuly updated",
	})

	return nil
}

func ListAllTasks(c *fiber.Ctx) error {
	User := c.Params("Token")
	Tasks, err:= Database.ListAllTasks(User)
	if err != nil {
		return c.Status(500).SendString("Task didn't listed")
	}

	c.JSON(fiber.Map{
		"tasks":Tasks,
	})

	return nil
}