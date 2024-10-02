package models

import "testing"

func TestNewTodo(t *testing.T) {
	expected := Todo{
		"test",
		false,
	}
	received := NewTodo("test")

	if received != expected {
		t.Errorf("expected %v, got %v", expected, received)
	}
}

func TestTodo_String(t *testing.T) {
	expected := "test - false"
	received := NewTodo("test").String()

	if received != expected {
		t.Errorf("expected %v, got %v", expected, received)
	}
}
