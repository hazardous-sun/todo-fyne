package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"todolist.com/models"
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

	data := []models.Todo{
		models.NewTodo("item 1"),
		models.NewTodo("item 2"),
		models.NewTodo("item 3"),
	}

	itemsList := widget.NewList(
		// function that returns the number of items in teh list
		func() int {
			return len(data)
		},
		// function that returns the component structure of the List Item
		func() fyne.CanvasObject {
			return container.NewBorder(
				nil, nil, nil,
				// "left" of the border
				widget.NewCheck("", func(b bool) {}),
				// takes the rest of the space
				widget.NewLabel(""),
			)
		},
		// function that is called for each item in the list and allows
		// you to show the content on the previously defined ui structure
		func(id widget.ListItemID, object fyne.CanvasObject) {
			ctr, _ := object.(*fyne.Container)
			lbl := ctr.Objects[0].(*widget.Label)
			check := ctr.Objects[1].(*widget.Check)
			lbl.SetText(data[id].Description)
			check.SetChecked(data[id].Done)
		},
	)

	w.SetContent(
		container.NewBorder(
			// TOP
			nil,
			// BOTTOM
			container.NewBorder(
				// inner - top
				nil,
				// inner - bottom
				nil,
				// inner - left
				nil,
				// inner - right
				addBtn,
				// inner - take the rest of the space
				newTodoDescTxt,
			),
			// LEFT
			nil,
			// RIGHT
			nil,
			// TAKE THE REST OF THE SPACE
			itemsList,
		),
	)
	w.ShowAndRun()
}
