package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"todolist.com/models"
)

func main() {
	a := app.New()
	w := a.NewWindow("TODO App")

	w.Resize(fyne.NewSize(300, 400))

	data := []models.Todo{
		models.LoadTodo("item 1", false),
		models.LoadTodo("item 2", true),
		models.LoadTodo("item 3", false),
	}

	todos := binding.NewUntypedList()

	for _, t := range data {
		err := todos.Append(t)

		if err != nil {
			panic(err)
		}
	}

	// new item description
	newTodoDescTxt := widget.NewEntry()
	newTodoDescTxt.PlaceHolder = "New TODO Description..."

	// add button
	addBtn := widget.NewButton("Add", func() {
		if len(newTodoDescTxt.Text) > 0 {
			todos.Append(models.NewTodo(newTodoDescTxt.Text))
			newTodoDescTxt.Text = ""
		}
	})
	addBtn.Disable()

	// basic validation for the description
	newTodoDescTxt.OnChanged = func(s string) {
		addBtn.Disable()

		if len(s) >= 3 {
			addBtn.Enable()
		}
	}

	//var selectedItem binding.DataItem

	itemsList := widget.NewListWithData(
		// the binding.List type
		todos,
		// function that returns the component structure of the List Item
		func() fyne.CanvasObject {
			lbl := widget.NewLabel("")
			checkbox := widget.NewCheck("", func(b bool) {
				if b {
					fmt.Println("item checked", lbl.Text)
				} else {
					fmt.Println("item not checked", lbl.Text)
				}
			})
			ctr := container.NewBorder(
				nil, nil, nil,
				// "left" of the border
				checkbox,
				// takes the rest of the space
				lbl,
			)
			return ctr
		},
		// function that is called for each item in the list and allows
		// you to show the content on the previously defined ui structure
		func(di binding.DataItem, object fyne.CanvasObject) {
			ctr, _ := object.(*fyne.Container)
			lbl := ctr.Objects[0].(*widget.Label)
			check := ctr.Objects[1].(*widget.Check)
			todo := NewTodoFromDataItem(di)
			lbl.SetText(todo.Description)
			check.SetChecked(todo.Done)

			// Mark the selected data item
			//selectedItem = di
			//fmt.Println(selectedItem)
			newTodoDescTxt.Text = todo.Description
		},
	)

	w.SetContent(
		container.NewBorder(
			// TOP
			nil,
			// BOTTOM
			container.NewBorder(
				nil, nil, nil,

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

func NewTodoFromDataItem(item binding.DataItem) models.Todo {
	v, _ := item.(binding.Untyped).Get()
	return v.(models.Todo)
}
