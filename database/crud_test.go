package database

import (
	"testing"
	"todolist.com/models"
)

func TestRead(t *testing.T) {
	client := InitializeClient()
	filter := models.NewFilters()
	expected := []models.Todo{
		models.Todo{
			"item 1",
			"test item 1",
			false,
		},
		models.Todo{
			"item 2",
			"test item 2",
			true,
		},
	}
	received, err := Read(client, filter, "todo_test")

	if err != nil {
		t.Errorf("Read returned an error %v", err)
	}

	for i, v := range received {
		if v != expected[i] {
			t.Errorf("expected %v, got %v", expected[i], v)
		}
	}
}
