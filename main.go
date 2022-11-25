package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/ura3107/blogapi/api"
)

func main() {
	c := mysql.Config{
		DBName:               os.Getenv("DB_NAME"),
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASSWORD"),
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	fmt.Println(c.FormatDSN())
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
