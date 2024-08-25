package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "root"
	dbPassword := "pass"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println("Error connecting to database:", err)
	} else {
		fmt.Println("Successfully connected to DB")
	}
}
