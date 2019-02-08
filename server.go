package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// var todos = make(map[string]Todo)
var todos []Todo

func todosHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("METHOD => ", req.Method)
	method := req.Method

	if method == "GET" {
		resp, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error na" + err.Error()))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(resp)
	}

	if method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "error : %v", err)
			return
		}

		var item Todo
		err = json.Unmarshal(body, &item)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("json unmarshal error" + err.Error()))
			return
		}

		todos = append(todos, item)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello " + method + " created todos"))
	}
}

func main() {
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todos", todosHandler)

	log.Fatal(http.ListenAndServe(":1234", nil))
}
