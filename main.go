package main

import (
	"fmt"
	"go-api/modules/database"
	"go-api/services"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	fmt.Println("Connected to database")

	// Get all todos
	// var todos []models.Todo
	// rows, err := database.Db.Query("SELECT id, activity_group_id, title, is_active, priority, create_at, update_at, IFNULL(delete_at, '') FROM todos")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// for rows.Next() {
	// 	todo := models.Todo{}
	// 	if err := rows.Scan(&todo.ID, &todo.Activity_Group_ID, &todo.Title, &todo.Is_Active, &todo.Priority, &todo.Create_At, &todo.Update_At, &todo.Delete_At); err != nil {
	// 		fmt.Println(err) // Exit if we get an error
	// 	}

	// 	// Append Employee to Employees
	// 	todos = append(todos, todo)

	// }

	// fmt.Println(todos)

	app := fiber.New()

	app.Get("/todos", services.GetAll)
	app.Get("todo/:id", services.GetTodo)
	app.Post("/todo", services.AddTodo)
	app.Delete("todo/:id", services.DeleteTodo)

	log.Fatal(app.Listen(":3000"))
}
