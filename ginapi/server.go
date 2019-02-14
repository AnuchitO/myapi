package main

import (
	"net/http"
	"strconv"

	"github.com/AnuchitO/myapi/todo"
	"github.com/gin-gonic/gin"
)

var todos = []todo.Todo{
	todo.Todo{ID: 0, Title: "homeworks", Status: "active"},
	todo.Todo{ID: 1, Title: "buy bmw", Status: "active"},
	todo.Todo{ID: 2, Title: "buy watch", Status: "completed"},
	todo.Todo{ID: 3, Title: "buy headphone", Status: "completed"},
}

var index = len(todos)

// SELECT items FROM TODO_TABLE WHERE status='active'
// reqest  /todos?status=active
func getTodosHandler(c *gin.Context) {
	status := c.Query("status")
	if status != "" {
		items := []todo.Todo{}
		for _, item := range todos {
			if item.Status == status {
				items = append(items, item)
			}
		}
		c.JSON(http.StatusOK, items)
		return
	}

	c.JSON(http.StatusOK, todos)
}

func createTodosHandler(c *gin.Context) {
	var item todo.Todo
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	index++
	item.ID = index
	item.Status = "active"

	todos = append(todos, item)

	c.String(http.StatusOK, "create todos successfull")
}

func getTodosByIdHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, item := range todos {
		if item.ID == id {
			c.JSON(http.StatusOK, item)
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{})
}

// localhost:1234/todos/:id => path param => c.Param
// localhost:1234/todos?status=active => query param => c.Query
func updateTodoHandler(c *gin.Context) {
	item := todo.Todo{}
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	for i, t := range todos {
		if t.ID == id {
			todos[i] = item
			c.JSON(http.StatusOK, gin.H{"status": "update success"})
			return
		}
	}

	c.JSON(http.StatusInternalServerError, gin.H{"status": "don't know"})
}

func deleteTodoHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tt := []todo.Todo{}
	for _, t := range todos {
		if t.ID != id {
			tt = append(tt, t)
		}
	}

	todos = tt

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func setUp() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api")

	v1.GET("/todos", getTodosHandler)
	v1.GET("/todos/:id", getTodosByIdHandler)
	v1.PUT("/todos/:id", updateTodoHandler)
	v1.DELETE("/todos/:id", deleteTodoHandler)
	v1.POST("/todos", createTodosHandler)

	return r
}
func main() {
	r := setUp()
	r.Run(":1234")
}
