// Package models :
// Contains the structs and functions for handling the filters used for selection in the database.
package models

import "fmt"

// Filters :
// Holds the data of which filters should be used
type Filters struct {
	Checked   bool
	Unchecked bool
}

// NewFilters :
// Initializes a Filters struct with all the flags set to false.
func NewFilters() Filters {
	return Filters{
		Checked:   false,
		Unchecked: false,
	}
}

// GetFlags :
// Returns an array of booleans representing each flag state.
func (f *Filters) GetFlags() []bool {
	var flags []bool
	flags = append(flags, f.Checked)
	flags = append(flags, f.Unchecked)
	return flags
}

// String :
// Returns the formatted Filters instance as a string.
func (f *Filters) String() string {
	return fmt.Sprintf("{%v \t %v}", f.Checked, f.Unchecked)
}
