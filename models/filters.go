package models

import "fmt"

type Filters struct {
	Checked   bool
	Unchecked bool
}

func NewFilters() Filters {
	return Filters{
		Checked:   true,
		Unchecked: true,
	}
}

func (f *Filters) GetFlags() []bool {
	var flags []bool
	flags = append(flags, f.Checked)
	flags = append(flags, f.Unchecked)
	return flags
}

func (f *Filters) String() string {
	return fmt.Sprintf("{%v \t %v}", f.Checked, f.Unchecked)
}
