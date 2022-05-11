package controllers

import (
	"go-api-gorm/database"
	"go-api-gorm/models"

	"github.com/gofiber/fiber/v2"
)

//AddActivity
func AddActivity(c *fiber.Ctx) error {
	activity := new(models.Activity)
	if err := c.BodyParser(activity); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBConn.Create(&activity)

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   activity,
	})
}

//GetActivityById
func GetActivityById(c *fiber.Ctx) error {
	activity := []models.Activity{}
	todos := []models.Todo{}

	database.DBConn.First(&activity, c.Params("id"))
	database.DBConn.Debug().Where("activity_group_id = ?", c.Params("id")).Find(&todos)

	if len(activity) > 0 {
		return c.Status(200).JSON(fiber.Map{
			"status":     "success",
			"data":       activity[0],
			"todo_items": todos,
		})
	} else {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "error",
			"error":  "Data Not Found",
		})
	}

}

//GetAllActivity
func GetAllActivity(c *fiber.Ctx) error {
	activity := []models.Activity{}

	database.DBConn.Find(&activity)

	return c.Status(200).JSON(fiber.Map{
		"status": "success",
		"data":   activity,
	})
}

//UpdateActivity
func UpdateActivity(c *fiber.Ctx) error {
	activity := []models.Activity{}
	data := new(models.Activity)
	if err := c.BodyParser(data); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBConn.Model(&activity).Where("id = ?", c.Params("id")).Update("title", data.Title)

	return c.Status(400).JSON(fiber.Map{
		"status":  "success",
		"message": "updated",
	})
}

//Delete
func DeleteActivity(c *fiber.Ctx) error {
	activity := []models.Activity{}
	// title := new(models.Activity)
	// if err := c.BodyParser(title); err != nil {
	// 	return c.Status(400).JSON(err.Error())
	// }
	if c.Params("id") == "" {
		return c.Status(400).JSON("id is required ")
	}
	database.DBConn.Where("id = ?", c.Params("id")).Delete(&activity)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "deleted",
	})
}
