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

	ctb := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT,
		status TEXT
	);
	`

	_, err = db.Exec(ctb)
	if err != nil {
		log.Fatal("can't create table", err)
	}

	fmt.Println("create table sucess")
}
