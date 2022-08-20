package controllers

//imports local and external dependencies required in this package
import (
	a "RESTAPI/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// struct for our controller
type ToDoController struct {
}

// function which is excecuted for the gettodo endpoint
func (toDocontroller *ToDoController) GetToDo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, a.ToDos)
}

// function which is excecuted for the createtodo endpoint
func (toDocontroller *ToDoController) CreateToDo(c *gin.Context) {
	//declare artist type
	var toDo a.ToDo

	// binds request body created variable and ensures no error
	if err := c.BindJSON(&toDo); err != nil {
		return
	}
	var toDos []a.ToDo = append(a.ToDos, toDo)
	c.IndentedJSON(http.StatusOK, toDos)
}
