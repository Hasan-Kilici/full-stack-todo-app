package main

import (
	"github.com/gofiber/fiber/v2"
	"todoList/Routers"
	"todoList/Database"
)

func main(){
	Database.Connect()
	app := fiber.New()

	Routers.Initalize(app)

	err := app.Listen("127.0.0.1:3000")
	if err != nil {
		panic(err)
	}
}