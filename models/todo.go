package models

import "fmt"

type Todo struct {
	Description string
	Done        bool
}

func NewTodo(description string) Todo {
	return Todo{description, false}
}

func LoadTodo(description string, checked bool) Todo {
	return Todo{
		description,
		checked,
	}
}

func (t Todo) String() string {
	return fmt.Sprintf("{%s \t %t}", t.Description, t.Done)
}
