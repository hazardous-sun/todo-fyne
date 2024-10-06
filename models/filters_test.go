package models

import "testing"

func TestNewFilters(t *testing.T) {
	expected := Filters{true, true}
	received := NewFilters()
	if received != expected {
		t.Errorf("expected %v, got %v", received, expected)
	}
}

func Test_GetFlags(t *testing.T) {
	expected := []bool{true, false}
	filtersStruct := Filters{true, false}
	received := filtersStruct.GetFlags()
	for i := 0; i < len(received); i++ {
		if received[i] != expected[i] {
			t.Errorf("expected %v, got %v", received, expected)
		}
	}
}

func TestTodo_String2(t *testing.T) {
	expected := "{true \t false}"
	received := Filters{true, false}

	if received.String() != expected {
		t.Errorf("expected %v, got %v", expected, received)
	}
}
