package models

// a struct type (collection of fields) to hold data
type ToDo struct {
	ID          string `json:"id"`
	Title       string `json:"title""`
	Description string `json:"description"`
}

// a slice with some mock data
var ToDos = []ToDo{
	{ID: "1", Title: "Barb", Description: "Barb"},
	{ID: "2", Title: "Clean", Description: "Clean"},
}
