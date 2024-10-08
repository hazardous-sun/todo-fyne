// Package models :
// Contains the struct and functions for handling the todo items.
package models

import "fmt"

// TodoItem :
// A representation of an item in the list of things to do.
type TodoItem struct {
	Title       string
	Description string
	Checked     bool
}

// LoadTodoItem :
// Builds an item to do in the list, that could be either true OR false.
func LoadTodoItem(title, description string, checked bool) TodoItem {
	return TodoItem{
		title,
		description,
		checked,
	}
}

// String :
// Returns the formatted TodoItem instance as a string.
func (t TodoItem) String() string {
	return fmt.Sprintf("{'%s' \t '%s' \t %t}", t.Title, t.Description, t.Checked)
}
