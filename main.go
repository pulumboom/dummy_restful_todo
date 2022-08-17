package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var list = []todo{
	{ID: "1", Item: "Clean Room", Completed: false},
	{ID: "2", Item: "Read Book", Completed: false},
	{ID: "3", Item: "Record Video", Completed: false},
}

func getList(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, list)
}

func addItem(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	list = append(list, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func toggleTodoStatus(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range list {
		if t.ID == id {
			return &list[i], nil
		}
	}
	return nil, errors.New("todo not found")
}

func getItem(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, todo)
}

func main() {
	router := gin.Default()
	router.GET("/list", getList)
	router.GET("/list/:id", getItem)
	router.PATCH("/list/:id", toggleTodoStatus)
	router.POST("/list", addItem)
	router.Run("localhost:8000")
}
