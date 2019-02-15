package todo

import (
	"log"
	"net/http"
	"strconv"

	"github.com/AnuchitO/myapi/database"
	"github.com/gin-gonic/gin"
)

func getTodosHandler(c *gin.Context) {
	status := c.Query("status")
	if status != "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "not support yet"})
		return
	}

	stmt, err := database.Conn().Prepare("SELECT id, title, status FROM todos")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't prepare query all todos statment" + err.Error()})
		return
	}

	rows, err := stmt.Query()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't query all todos" + err.Error()})
		return
	}

	var todos = []Todo{}

	for rows.Next() {
		t := Todo{}
		err := rows.Scan(&t.ID, &t.Title, &t.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "can't Scan row into variable" + err.Error()})
			return
		}

		todos = append(todos, t)
	}

	c.JSON(http.StatusOK, todos)
}

func createTodosHandler(c *gin.Context) {
	var item Todo
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	row := database.InsertTodo(item.Title, item.Status)
	err = row.Scan(&item.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't Scan row into variable" + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}

func getTodoByIdHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// TODO: handle error

	stmt, err := database.Conn().Prepare("SELECT id, title, status FROM todos WHERE id=$1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	row := stmt.QueryRow(id)

	t := Todo{}
	err = row.Scan(&t.ID, &t.Title, &t.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "data not found"})
		return
	}

	c.JSON(http.StatusOK, t)
}

func updateTodoHandler(c *gin.Context) {
	item := Todo{}
	err := c.ShouldBindJSON(&item)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	stmt, err := database.Conn().Prepare("UPDATE todos SET status=$2 WHERE id=$1;")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	if _, err := stmt.Exec(id, item.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	item.ID = id

	c.JSON(http.StatusOK, item)
}

func deleteTodoHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	// TODO: handle error

	stmt, err := database.Conn().Prepare("DELETE FROM todos WHERE id = $1")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if _, err := stmt.Exec(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api")

	v1.GET("/todos", getTodosHandler)
	v1.GET("/todos/:id", getTodoByIdHandler)
	v1.PUT("/todos/:id", updateTodoHandler)
	v1.DELETE("/todos/:id", deleteTodoHandler)
	v1.POST("/todos", createTodosHandler)

	return r
}

func CreateTable() {
	ctb := `
	CREATE TABLE IF NOT EXISTS todos (
		id SERIAL PRIMARY KEY,
		title TEXT,
		status TEXT
	);`

	_, err := database.Conn().Exec(ctb)
	if err != nil {
		log.Fatal("can't create table", err)
	}
}
