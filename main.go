package main

import (
	"log"

	"go-api-gorm/controllers"
	"go-api-gorm/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/hello", controllers.Hello)
	app.Get("/allbooks", controllers.AllBooks)
	app.Get("/book/:id", controllers.GetBook)
	app.Post("/book", controllers.AddBook)
	app.Put("/book", controllers.Update)
	app.Delete("/book", controllers.Delete)
}

func ActivityRoutes(app *fiber.App) {
	app.Get("/activity-groups", controllers.GetAllActivity)
	app.Get("/activity-groups/:id", controllers.GetActivityById)
	app.Post("/activity-groups", controllers.AddActivity)
	app.Put("/activity-groups/:id", controllers.UpdateActivity)
	app.Delete("/activity-groups/:id", controllers.DeleteActivity)
}

func TodoRouter(app *fiber.App) {
	app.Get("/todos", controllers.GetAllTodo)
	app.Get("/todo/:id", controllers.GetTodoById)
	app.Post("/todo", controllers.AddTodo)
	app.Put("/todo/:id", controllers.UpdateTodo)
	app.Delete("/todo/:id", controllers.DeleteTodo)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	// fix bug cors vue axios golang
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		return c.Next()
	})

	setUpRoutes(app)
	ActivityRoutes(app)
	TodoRouter(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3030"))
}
