package toDo

type ToDo struct {
	ID          string `json:"id"`
	Title       string `json:"title""`
	Description string `json:"description"`
}

var ToDos = []ToDo{
	{ID: "1", Title: "Barb", Description: "Barb"},
	{ID: "2", Title: "Clean", Description: "Clean"},
}
