package models

import "fmt"

type Todo struct {
	Description string
	Done        bool
}

func NewTodo(description string) Todo {
	return Todo{description, false}
}

func (t Todo) String() string {
	return fmt.Sprintf("%s - %t", t.Description, t.Done)
}
