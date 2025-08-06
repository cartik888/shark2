package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, name)

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic("Database connection failed: " + err.Error())
	}

	err = DB.Ping()
	if err != nil {
		panic("DB ping failed: " + err.Error())
	}

	fmt.Println("✅ Connected to GCP MySQL successfully!")
}
