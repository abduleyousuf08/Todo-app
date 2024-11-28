package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


type Todo struct {
	ID        int    `json:"id"`          
	Title     string `json:"title"`      
	Completed bool   `json:"completed"`  
}

var todos = []Todo{
	{ID: 1, Title: "Buy milk", Completed: false},
	{ID: 2, Title: "Build todo app", Completed: false},
	{ID: 3, Title: "Clean the house", Completed: false},
}


func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}


func CreateTodo(context *gin.Context){
	var newTodo Todo

	if err := context.BindJSON(&newTodo); err != nil{
		return
	}

	todos = append(todos,newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func GetTodoById(id int) (*Todo, error) {
	for i,v := range todos{
		if v.ID == id{
			return &todos[i], nil
		}

	}
	return nil, errors.New("todo not found")
}

func GetTodo(context *gin.Context){
	idParam := context.Param("id")

	//Converting the id to int
	id, err := strconv.Atoi(idParam)
	
    if err != nil {

    fmt.Println("Error converting id:", err)
	}

	todo, err := GetTodoById(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"Error" : "todo not found"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"Here is the todo you've searched":todo})
}


func UpdateTodo(context *gin.Context){
	idParam := context.Param("id")

	//Converting the id to int
	id, err := strconv.Atoi(idParam)
	
    if err != nil {

    fmt.Println("Error converting id:", err)
	}

	todo, err := GetTodoById(id)

	if err != nil{
		context.IndentedJSON(http.StatusNotFound, gin.H{"Error" : "todo not found"})
		return
	}

	todo.Completed = !todo.Completed
	context.IndentedJSON(http.StatusOK, todo)
}

func DeleteTodo(context *gin.Context) {
	idParam := context.Param("id")

	// Convert id to int
	id, err := strconv.Atoi(idParam)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Check if the todo exists
	_, err = GetTodoById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// Filter out the todo with the given id
	newTodos := []Todo{} // Create a new slice to hold todos except the one being deleted
	for _, v := range todos {
		if v.ID != id {
			newTodos = append(newTodos, v)
		}
	}

	// Update the todos list
	todos = newTodos

	// Respond with the updated todos list
	context.IndentedJSON(http.StatusOK, gin.H{
		"message": "Todo deleted successfully",
		"todos":   todos,
	})
}


func main() {
	router := gin.Default()
	router.GET("/todos", GetTodos)
	router.GET("/todos/:id", GetTodo)
	router.POST("/todos", CreateTodo)
	router.DELETE("/todos/:id",DeleteTodo )
	router.PUT("/todos/:id", UpdateTodo)

	router.Run("localhost:9090")
}
