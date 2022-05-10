package main

import (
	"log"

	"go-api-gorm/database"
	"go-api-gorm/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func setUpRoutes(app *fiber.App) {
	app.Get("/hello", routes.Hello)
	app.Get("/allbooks", routes.AllBooks)
	app.Get("/book/:id", routes.GetBook)
	app.Post("/book", routes.AddBook)
	app.Put("/book", routes.Update)
	app.Delete("/book", routes.Delete)
}

func ActivityRoutes(app *fiber.App) {
	app.Get("/activity", routes.GetAllActivity)
	app.Get("/activity/:id", routes.GetActivityById)
	app.Post("/activity", routes.AddActivity)
	app.Put("/activity/:id", routes.UpdateActivity)
	app.Delete("/activity/:id", routes.DeleteActivity)
}

func main() {
	database.ConnectDb()
	app := fiber.New()

	setUpRoutes(app)
	ActivityRoutes(app)

	app.Use(cors.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})

	log.Fatal(app.Listen(":3030"))
}
