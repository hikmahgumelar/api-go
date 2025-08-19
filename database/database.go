package database

import (
  "database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"

)

func SetupDB() *sql.DB {
    db, err := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/api-go?parseTime=true")
		if err != nil {
			log.Fatal(err)
		}
    err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("database Terkoneksi")
		return db 

}
