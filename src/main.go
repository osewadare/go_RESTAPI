package main

import (
	artistController "RESTAPI/controllers/toDocontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/todos", artistController.GetToDo)
	router.GET("/createtodos", artistController.CreateToDo)
	router.Run("localhost:8080")
}
