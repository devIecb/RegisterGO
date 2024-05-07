package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	if err != nil{
		log.Println("Error to load .env", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS") 
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", connStr); 
	if err != nil{
		log.Println("Error to connect in DB", err)
	}
	return db, nil
}