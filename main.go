package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/YamaguchiKoki/go_prc/api"
)

var (
	dbUser = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbHost = "db"
	dbConn = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("failed to connect db")
		return
	}

	r := api.NewRouter(db)

	log.Println("server start at port 8080")

	//log.Fatal: 実行時異常としてプログラム終了
	log.Fatal(http.ListenAndServe(":8080", r))
}