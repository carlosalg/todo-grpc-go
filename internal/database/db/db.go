package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "todoapp_db",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingError := db.Ping()
	if pingError != nil {
		log.Fatal(pingError)
	}
	fmt.Println("Connected!")
	DB = db
}

func CloseDB() {
	if DB != nil {
		DB.Close()
		fmt.Println("Database connection closed.")
	}
}
