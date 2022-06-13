package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var database sql.DB
var databaseName string

func Connect() error {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	databaseName = os.Getenv("DB_DATABASE_NAME")

	dbconf := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + databaseName + "?charset=utf8mb4"
	db, err := sql.Open("mysql", dbconf)
	database = *db
	return err
}

func Close() {
	database.Close()
}

func Ping() error {
	return database.Ping()
}

func TableList() (*sql.Rows, error) {
	res, err := database.Query("SHOW TABLES FROM " + databaseName)
	if err != nil {
		return nil, err
	} else {
		return res, err
	}
}
