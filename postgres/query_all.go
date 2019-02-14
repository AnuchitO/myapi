package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/AnuchitO/myapi/todo"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Connect to database error", err)
	}
	defer db.Close()

	stmt, err := db.Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		log.Fatal("can't prepare query all todos statment", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal("can't query all todos", err)
	}

	var todos = []todo.Todo{}

	for rows.Next() {
		t := todo.Todo{}
		err := rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			log.Fatal("can't Scan row into variable", err)
		}

		todos = append(todos, t)
	}

	fmt.Printf("% #v\n", todos)
	fmt.Println("query all todos success")
}
