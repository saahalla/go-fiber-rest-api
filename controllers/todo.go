package controllers

import (
	"go-api-gorm/database"
	"go-api-gorm/models"

	"github.com/gofiber/fiber/v2"
)

//AddTodo
func AddTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)
	if err := c.BodyParser(todo); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if todo.Priority == "" {
		todo.Priority = "very-high"
	}
	todo.IsActive = "true"

	database.DBConn.Create(&todo)

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   todo,
	})
}

//GetTodoId
func GetTodoById(c *fiber.Ctx) error {
	todo := []models.Todo{}

	database.DBConn.First(&todo, c.Params("id"))

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   todo,
	})
}

//GetAllTodo
func GetAllTodo(c *fiber.Ctx) error {
	todo := []models.Todo{}

	database.DBConn.Find(&todo)

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   todo,
	})
}

//UpdateTodo
func UpdateTodo(c *fiber.Ctx) error {
	todo := []models.Todo{}
	data := new(models.Todo)
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBConn.Model(&todo).Where("id = ?", c.Params("id")).Update("title", data.Title)

	return c.Status(400).JSON(fiber.Map{
		"status":  "success",
		"message": "updated",
	})
}

//DeleteTodo
func DeleteTodo(c *fiber.Ctx) error {
	todo := []models.Todo{}
	// title := new(models.Todo)
	// if err := c.BodyParser(title); err != nil {
	// 	return c.Status(400).JSON(err.Error())
	// }
	if c.Params("id") == "" {
		return c.Status(400).JSON("id is required ")
	}
	database.DBConn.Where("id = ?", c.Params("id")).Delete(&todo)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "deleted",
	})
}
