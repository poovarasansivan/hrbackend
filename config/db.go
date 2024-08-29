package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var Database *sql.DB
var err error

func ConnectDB() {
	dsn := "uodjdfoophd8wtcx:LhCTdiGCojRoKW9vcsyC@tcp(bamuzjys35fimiyhi0zw-mysql.services.clever-cloud.com:3306)/bamuzjys35fimiyhi0zw"
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
