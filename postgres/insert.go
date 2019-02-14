package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("can't connect to database ", err)
	}
	defer db.Close()

	row := db.QueryRow("INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING id", "buy shoes", "active")
	var id int
	err = row.Scan(&id)
	if err != nil {
		log.Fatal("can't scan id", err)
	}
	fmt.Println("insert to success id : ", id)
}
