package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

const (
	host     = "localhost"
	port     = 3306
	user     = "saahalla"
	password = "sahal07seven"
	dbname   = "db_devcode"
)

func Connect() error {
	var err error

	Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		return err
	}
	if err = Db.Ping(); err != nil {
		return err
	}
	fmt.Println("Connected to database")
	return nil
}
