// Package models :
// Contains the struct and functions for handling the todo items.
package models

import "fmt"

// Todo :
// A representation of an item in the list of things to do.
type Todo struct {
	Description string
	Checked     bool
}

// NewTodo :
// Initializes an unchecked item in the list.
func NewTodo(description string) Todo {
	return Todo{description, false}
}

// LoadTodo :
// Builds an item to do in the list, that could be either true OR false.
func LoadTodo(description string, checked bool) Todo {
	return Todo{
		description,
		checked,
	}
}

// String :
// Returns the formatted Todo instance as a string.
func (t Todo) String() string {
	return fmt.Sprintf("{%s \t %t}", t.Description, t.Checked)
}
