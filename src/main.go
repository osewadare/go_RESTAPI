package main

import (
	controller "RESTAPI/controllers"
	"github.com/gin-gonic/gin"
)

// Entry point for the Go Program
func main() {

	//creates router with default settings
	router := gin.Default()
	toDoController := controller.ToDoController{}

	//sets up endpoints for gettodos and createtodo and a server listening on port 8080
	router.GET("/todos", toDoController.GetToDo)
	router.GET("/createtodos", toDoController.CreateToDo)
	router.Run("localhost:8080")
}
