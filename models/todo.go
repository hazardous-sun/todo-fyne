// Package models :
// Contains the struct and functions for handling the todo items.
package models

import "fmt"

// Todo :
// A representation of an item in the list of things to do.
type Todo struct {
	Title       string
	Description string
	Checked     bool
}

// LoadTodo :
// Builds an item to do in the list, that could be either true OR false.
func LoadTodo(title, description string, checked bool) Todo {
	return Todo{
		title,
		description,
		checked,
	}
}

// String :
// Returns the formatted Todo instance as a string.
func (t Todo) String() string {
	return fmt.Sprintf("{'%s' \t '%s' \t %t}", t.Title, t.Description, t.Checked)
}
