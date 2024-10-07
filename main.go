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

const MinimumDescriptionLen = 3

func main() {
	// Connect to the database client
	dbClient = database.InitializeClient()

	// Initialize filters
	filters = models.NewFilters()

	// Run the GUI application
	run()
}

func run() {
	// initialize the app and the initial window for the application
	a := app.New()
	w := a.NewWindow("TODO App")

	w.Resize(fyne.NewSize(300, 400))

	// collect the data from the DB
	todos := getTodos()

	// container with the checkboxes for filtering the items
	filtersCtr := initializeFilterCtr()

	// new item description
	newItemEntry := initializeNewItemEntry()

	// add button
	addBtn := initializeAddBtn(newItemEntry, todos)

	// del button
	delBtn := initializeDelBtn(newItemEntry)

	// basic validation for the description
	newItemEntry.OnChanged = func(s string) {
		addBtn.Disable()
		delBtn.Disable()

		if len(s) >= MinimumDescriptionLen {
			addBtn.Enable()
			delBtn.Enable()
		}
	}

	// list that holds the items to do
	itemsList = initializeItemsList(todos)

	// pass the values to the window
	w.SetContent(
		container.NewBorder(
			// TOP
			filtersCtr,
			// BOTTOM
			container.NewBorder(
				nil, nil,

				// inner - left
				delBtn,

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

// Reads from the database and returns an untyped list used to store the values of the items to do.
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

// Receives a data item from an untyped list and casts it to a models.Todo instance.
func newTodoFromDataItem(item binding.DataItem) models.Todo {
	v, _ := item.(binding.Untyped).Get()
	return v.(models.Todo)
}

// Iterates over a bindings.UntypedList and returns the index that contains a models.Todo instance with the same
// description of the string parameter passed.
func getTodoFromList(lbl string) int {
	for i := 0; i < todos.Length(); i++ {
		di, err := todos.GetItem(i)

		if err != nil {
			panic(err)
		}

		todo := newTodoFromDataItem(di)

		if todo.Description == lbl {
			return i
		}
	}
	return -1
}

// widgets -------------------------------------------------------------------------------------------------------------

// Initializes the container that holds the checkboxes for filtering the content displayed.
func initializeFilterCtr() *fyne.Container {
	return container.NewBorder( // TODO fix the filter functions
		nil, nil,
		widget.NewCheck(
			"Checked items",
			func(b bool) {
				filters.Checked = b
				fmt.Printf("Checked items = %v \n", b)
			},
		),
		widget.NewCheck(
			"Unchecked items",
			func(b bool) {
				filters.Unchecked = b
				fmt.Printf("Unchecked items = %v \n", b)
			},
		),
	)
}

// Initializes the button used for adding new elements to the list.
func initializeAddBtn(newItemEntry *widget.Entry, todos binding.UntypedList) *widget.Button {
	addBtn := widget.NewButton("Add", func() {
		if getTodoFromList(newItemEntry.Text) == -1 {
			err := todos.Append(models.NewTodo(newItemEntry.Text))
			database.Create(dbClient, newItemEntry.Text, false)

			if err != nil {
				return
			}

			newItemEntry.Text = ""
		}
	})
	return addBtn
}

// Initializes the button used for removing elements from the list.
func initializeDelBtn(newItemEntry *widget.Entry) *widget.Button {
	delBtn := widget.NewButton("Delete", func() {
		index := getTodoFromList(newItemEntry.Text)
		if index != -1 {
			fmt.Println("O ITEM EXISTE NA LISTA")
			di, _ := todos.GetItem(index)
			todo := newTodoFromDataItem(di)
			err := todos.Remove(todo)

			if err != nil {
				panic(err)
			}

			database.Delete(dbClient, todo.Description)
		} else {
			fmt.Println("Item does not exist in the list.")
		}
	})
	return delBtn
}

// Initializes the text entry for typing the new item description.
func initializeNewItemEntry() *widget.Entry {
	newItemEntry := widget.NewEntry()
	newItemEntry.PlaceHolder = "New TODO Description..."
	return newItemEntry
}

// Initializes the list used for displaying the items to do.
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

// Initializes the checkboxes used for showing if the item is completed or not.
func initializeCheckbox(lbl *widget.Label) *widget.Check {
	return widget.NewCheck("", func(b bool) {
		index := getTodoFromList(lbl.Text)
		err := todos.SetValue(index, models.LoadTodo(
			lbl.Text,
			b,
		))

		if err != nil {
			panic(err)
		}

		database.Update(dbClient, lbl.Text, b)
	})
}
