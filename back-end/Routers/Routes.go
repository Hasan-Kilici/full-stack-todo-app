package Routers

import(
	"todoList/Middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"todoList/Handlers"
)

func Initalize(app *fiber.App){
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders:  "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
	
	app.Get("/", Handlers.HelloHandler)
	app.Get("/list/tasks/:Token", Handlers.ListAllTasks)

	app.Post("/create/task", Handlers.CreateTask)

	app.Delete("/delete/task/:Token", Handlers.DeleteTask)
	app.Delete("/delete/all/tasks", Handlers.DeleteAllTasks)

	app.Put("/update/task/:Token", Handlers.UpdateTaskStatus)

	app.Use(Middleware.Security)
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}