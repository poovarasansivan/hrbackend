package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Database *sql.DB
var err error

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/hrportal"
	Database, err = sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error connecting to the database: ", err)
		return
	}
	if err = Database.Ping(); err != nil {
		fmt.Println("Error pinging the database: ", err)
		return
	}
	fmt.Println("Database connected successfully!")
}

func GetDB() *sql.DB {
	return Database
}
