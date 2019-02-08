package main

import (
	"net/http"
	"strconv"

	"github.com/AnuchitO/myapi/todo"
	"github.com/gin-gonic/gin"
)

var todos = []todo.Todo{
	todo.Todo{ID: "1", Title: "buy cars", Status: "active"},
}

var id = 0

func getTodosHandler(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func createTodosHandler(c *gin.Context) {
	var item todo.Todo
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id = id + 1
	item.ID = strconv.Itoa(id)
	item.Status = "active"

	todos = append(todos, item)

	c.String(http.StatusOK, "create todos successfull")
}

func main() {
	r := gin.Default()
	r.GET("/todos", getTodosHandler)
	r.POST("/todos", createTodosHandler)
	r.Run(":1234")
}
