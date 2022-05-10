package database

import (
	"fmt"
	"go-api-gorm/config"
	"log"
	"os"

	"go-api-gorm/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

// connectDb
func ConnectDb() {
	config := map[string]string{
		"host":     config.GetEnv("DB_HOST"),
		"username": config.GetEnv("DB_USERNAME"),
		"password": config.GetEnv("DB_PASSWORD"),
		"dbname":   config.GetEnv("DB_NAME"),
		"port":     config.GetEnv("DB_PORT"),
	}

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config["username"], config["password"], config["host"], config["port"], config["dbname"])
	/*
		NOTE:
		To handle time.Time correctly, you need to include parseTime as a parameter. (more parameters)
		To fully support UTF-8 encoding, you need to change charset=utf8 to charset=utf8mb4. See this article for a detailed explanation
	*/

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.Activity{})
	DBConn = db

}
