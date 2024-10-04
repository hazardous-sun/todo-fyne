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
	// initialize test items
	todos := initializeTestTodos()

	// new item description
	newItemEntry := initializeNewItemEntry()

	// add button
	addBtn := initializeAddBtn(newItemEntry, todos)

	// basic validation for the description
	newItemEntry.OnChanged = func(s string) {
		addBtn.Disable()

		if len(s) >= 3 {
			addBtn.Enable()
		}
	}

	itemsList := initializeItemsList(todos, newItemEntry)

	run(addBtn, newItemEntry, itemsList)
}

func initializeTestTodos() binding.UntypedList {
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

	return todos
}

func initializeAddBtn(newItemEntry *widget.Entry, todos binding.UntypedList) *widget.Button {
	addBtn := widget.NewButton("Add", func() {
		if len(newItemEntry.Text) > 0 {
			todos.Append(models.NewTodo(newItemEntry.Text))
			newItemEntry.Text = ""
		}
	})
	addBtn.Disable()
	return addBtn
}

func initializeNewItemEntry() *widget.Entry {
	newItemEntry := widget.NewEntry()
	newItemEntry.PlaceHolder = "New TODO Description..."
	return newItemEntry
}

func initializeItemsList(todos binding.UntypedList, newItemEntry *widget.Entry) *widget.List {
	return widget.NewListWithData(
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
			todo := newTodoFromDataItem(di)
			lbl.SetText(todo.Description)
			check.SetChecked(todo.Done)

			// Mark the selected data item
			newItemEntry.Text = todo.Description
		},
	)
}

func run(btn *widget.Button, entry *widget.Entry, list *widget.List) {
	a := app.New()
	w := a.NewWindow("TODO App")

	w.Resize(fyne.NewSize(300, 400))

	w.SetContent(
		container.NewBorder(
			// TOP
			nil,
			// BOTTOM
			container.NewBorder(
				nil, nil, nil,

				// inner - right
				btn,
				// inner - take the rest of the space
				entry,
			),
			// LEFT
			nil,
			// RIGHT
			nil,
			// TAKE THE REST OF THE SPACE
			list,
		),
	)
	w.ShowAndRun()
}

func newTodoFromDataItem(item binding.DataItem) models.Todo {
	v, _ := item.(binding.Untyped).Get()
	return v.(models.Todo)
}
