package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Conn() *sql.DB {
	if db != nil {
		fmt.Println("olddddddddddd")
		return db
	}

	var err error
	db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("can't connect to database", err)
	}

	fmt.Println("nwwwwwwwwwww")
	return db
}

func InsertTodo(title, status string) *sql.Row {
	return Conn().QueryRow("INSERT INTO todos (title, status) VALUES ($1, $2) RETURNING id", title, status)
}
