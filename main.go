package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/supabase-community/supabase-go"
	"todolist.com/database"
	"todolist.com/models"
)

var dbClient *supabase.Client
var filters models.Filters
var itemsList *widget.List
var todos binding.UntypedList

func main() {
	// Connect to the database client
	dbClient = database.InitializeClient()

	// Initialize filters
	filters = models.NewFilters()

	// Run the GUI application
	run()
}

func run() {
	todos := getTodos()

	// container with the checkboxes for filtering the items
	filtersCtr := initializeFilterCtr(&todos)

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

	// list that holds the items to do
	itemsList = initializeItemsList(todos)

	a := app.New()
	w := a.NewWindow("TODO App")

	w.Resize(fyne.NewSize(300, 400))

	w.SetContent(
		container.NewBorder(
			// TOP
			filtersCtr,
			// BOTTOM
			container.NewBorder(
				nil, nil, nil,

				// inner - right
				addBtn,
				// inner - take the rest of the space
				newItemEntry,
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

// data parsing --------------------------------------------------------------------------------------------------------

func getTodos() binding.UntypedList {
	// Collect the existing values inside the database
	data := database.Read(dbClient, filters)

	// Initialize list
	todos = binding.NewUntypedList()

	// Append values to list
	for _, t := range data {
		err := todos.Append(t)

		if err != nil {
			panic(err)
		}
	}

	return todos
}

func newTodoFromDataItem(item binding.DataItem) models.Todo {
	v, _ := item.(binding.Untyped).Get()
	return v.(models.Todo)
}

// widgets -------------------------------------------------------------------------------------------------------------

func initializeFilterCtr(todos *binding.UntypedList) *fyne.Container {
	return container.NewBorder(
		nil, nil,
		widget.NewCheck(
			"Checked items",
			func(b bool) {
				filters.Checked = b
				*todos = getTodos()
			},
		),
		widget.NewCheck(
			"Unchecked items",
			func(b bool) {
				filters.Unchecked = b
				*todos = getTodos()
			},
		),
	)
}

func initializeAddBtn(newItemEntry *widget.Entry, todos binding.UntypedList) *widget.Button {
	addBtn := widget.NewButton("Add", func() {
		if len(newItemEntry.Text) > 0 {
			err := todos.Append(models.NewTodo(newItemEntry.Text)) // TODO add a call to Create to send the new item to the DB
			database.Create(dbClient, newItemEntry.Text, false)

			if err != nil {
				return
			}

			newItemEntry.Text = ""
		}
	})
	return addBtn
}

func initializeNewItemEntry() *widget.Entry {
	newItemEntry := widget.NewEntry()
	newItemEntry.PlaceHolder = "New TODO Description..."
	return newItemEntry
}

func initializeItemsList(todos binding.UntypedList) *widget.List {
	return widget.NewListWithData(
		// the binding.List type
		todos,
		// function that returns the component structure of the List Item
		func() fyne.CanvasObject {
			lbl := widget.NewLabel("")
			checkbox := initializeCheckbox(lbl)
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
			check.SetChecked(todo.Checked)
		},
	)
}

func initializeCheckbox(lbl *widget.Label) *widget.Check {
	return widget.NewCheck("", func(b bool) {
		if b {
			fmt.Println("item checked", lbl.Text)
		} else {
			fmt.Println("item not checked", lbl.Text)
		}
	})
}

// ---------------------------------------------------------------------------------------------------------------------
