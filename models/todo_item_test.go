package models

import "testing"

func TestLoadTodo(t *testing.T) {
	expected := TodoItem{
		"test",
		"test description",
		true,
	}
	received := LoadTodoItem("test", "test description", true)
	if received != expected {
		t.Errorf("expected %v, got %v", expected, received)
	}
}

func TestTodo_String(t *testing.T) {
	expected := "{'test' \t 'test description' \t false}"
	received := LoadTodoItem("test", "test description", false).String()

	if received != expected {
		t.Errorf("expected %v, got %v", expected, received)
	}
}
