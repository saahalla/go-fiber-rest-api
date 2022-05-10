package services

import (
	"fmt"
	"go-api/models"
	"go-api/modules/database"
	"log"

	"github.com/gofiber/fiber/v2"
)

func GetAll(c *fiber.Ctx) error {
	var todos []models.Todo
	rows, err := database.Db.Query("SELECT id, activity_group_id, title, is_active, priority, create_at, update_at, IFNULL(delete_at, '') FROM todos")

	if err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	for rows.Next() {
		todo := models.Todo{}
		if err := rows.Scan(&todo.ID, &todo.Activity_Group_ID, &todo.Title, &todo.Is_Active, &todo.Priority, &todo.Create_At, &todo.Update_At, &todo.Delete_At); err != nil {
			return c.JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}

		// Append Employee to Employees
		todos = append(todos, todo)

	}

	fmt.Println(todos)
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   todos,
	})
}

func GetTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	var todo models.Todo
	rows, err := database.Db.Query("SELECT * FROM todos WHERE id = ?", id)

	if err != nil {
		fmt.Println(err)
		return c.JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	for rows.Next() {
		if err := rows.Scan(&todo.ID, &todo.Activity_Group_ID, &todo.Title, &todo.Is_Active, &todo.Priority, &todo.Create_At, &todo.Update_At, &todo.Delete_At); err != nil {
			fmt.Println(err.Error())
			return c.JSON(fiber.Map{
				"status": "error",
				"error":  err.Error(),
			})
		}
	}
	// if data not found
	if todo.ID == 0 {
		fmt.Println(todo)
		return c.JSON(fiber.Map{
			"status": "error",
			"error":  "data not found",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   todo,
	})
}

func AddTodo(c *fiber.Ctx) error {
	newtodo := new(models.Todo)

	if err := c.BodyParser(newtodo); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}

	// insert todo
	res, err := database.Db.Query("INSERT INTO todos (activity_group_id, title, create_at, update_at) VALUES (?, ?, now(), now())", newtodo.Activity_Group_ID, newtodo.Title)
	if err != nil {
		return err
	}
	defer res.Close()

	// rows, err := database.Db.Query("SELECT * from todos WHERE id = LAST_INSERT_ID()")
	fmt.Println(res)
	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"title": newtodo.Title,
		},
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	res, err := database.Db.Query("DELETE FROM todos WHERE id = ?", id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status": "error",
			"error":  err.Error(),
		})
	}
	log.Println(res)
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Data Successfully Deleted",
	})
}
