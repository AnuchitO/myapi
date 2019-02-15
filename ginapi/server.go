package main

import (
	"github.com/AnuchitO/myapi/todo"
)

func main() {
	todo.CreateTable()

	r := todo.NewRouter()
	r.Run(":1234")
}
