package toDoController

import (
	a "RESTAPI/models/toDo"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetToDo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, a.ToDos)
}

func CreateToDo(c *gin.Context) {
	//declare artist type
	var newAlbum a.ToDo

	// binds request body to newalbum and ensures no error
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}
	var albums []a.ToDo = append(a.ToDos, newAlbum)
	c.IndentedJSON(http.StatusOK, albums)
}
