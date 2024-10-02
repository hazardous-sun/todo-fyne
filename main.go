package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("TODO App")

	w.Resize(fyne.NewSize(300, 400))

	newTodoDescTxt := widget.NewEntry()
	newTodoDescTxt.PlaceHolder = "New TODO Description..."
	addBtn := widget.NewButton("Add", func() {
		fmt.Println(newTodoDescTxt.Text + " needs to be added to the TODO list!")
	})
	addBtn.Disable()

	newTodoDescTxt.OnChanged = func(s string) {
		addBtn.Disable()

		if len(s) >= 3 {
			addBtn.Enable()
		}
	}

	w.SetContent(
		container.NewBorder(
			nil, // TOP of the container

			container.NewBorder(
				nil, // TOP
				nil, // BOTTOM
				nil, // Left
				// RIGHT ↓
				addBtn,
				// take the rest of the space
				newTodoDescTxt,
			),

			nil,
			nil,
		),
	)
	w.ShowAndRun()
}
