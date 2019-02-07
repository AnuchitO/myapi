package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("hello handler")
	hello := []byte(`{"name": "anuchito"}`)
	w.Write(hello)
}

func main() {
	fmt.Println("start...")

	http.HandleFunc("/", helloHandler)
	http.ListenAndServe(":1234", nil)

	fmt.Println("end!")
}
