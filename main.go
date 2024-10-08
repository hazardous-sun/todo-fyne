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
var appInstance fyne.App

const MinimumDescriptionLen = 3

func main() {
	// Connect to the database client
	dbClient = database.InitializeClient()

	if dbClient == nil {
		panic(fmt.Errorf("failed to initialize db client"))
	}

	// initialize the app
	appInstance = app.New()

	// Initialize filters
	filters = models.NewFilters()

	// Run the GUI application
	run()
}

// Loading screen ------------------------------------------------------------------------------------------------------

func initializeLoadingScreen() fyne.Window {
	w := appInstance.NewWindow("Loading app")
	w.Resize(fyne.NewSize(300, 400))
	loadingLayout := container.NewBorder(
		widget.NewLabel("Connecting to the database..."),
		widget.NewProgressBar(),
		nil, nil,
	)
	w.SetContent(loadingLayout)
	return w
}

// ---------------------------------------------------------------------------------------------------------------------

func run() {
	// initialize the window with items to do
	w := appInstance.NewWindow("TODO App")

	w.Resize(fyne.NewSize(300, 400))

	// collect the data from the DB
	todos, err := getTodos()

	if err != nil {
		panic(err)
	}

	// container with the checkboxes for filtering the items
	filtersCtr := initializeFilterCtr()

	// add button
	addBtn := initializeAddBtn(w)

	// del button
	//delBtn := initializeDelBtn(newItemEntry)

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
				//delBtn,
				nil,

				// inner - right
				nil,
				addBtn,
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
func getTodos() (binding.UntypedList, error) {
	// Collect the existing values inside the database
	data, err := database.Read(dbClient, filters, "todo")

	if err != nil {
		return nil, err
	}

	// Initialize list
	todos = binding.NewUntypedList()

	// Append values to list
	for _, t := range data {
		err := todos.Append(t)

		if err != nil {
			panic(err)
		}
	}

	return todos, nil
}

// Receives appInstance data item from an untyped list and casts it to appInstance models.Todo instance.
func newTodoFromDataItem(item binding.DataItem) models.Todo {
	v, _ := item.(binding.Untyped).Get()
	return v.(models.Todo)
}

// Iterates over appInstance bindings.UntypedList and returns the index that contains appInstance models.Todo instance with the same
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
			},
		),
		widget.NewCheck(
			"Unchecked items",
			func(b bool) {
				filters.Unchecked = b
			},
		),
	)
}

// Initializes the button used for adding new elements to the list.
func initializeAddBtn(origin fyne.Window) *widget.Button {
	btn := widget.NewButton("Add", func() {
		// initialize the new window
		w := appInstance.NewWindow("Add item")
		w.Resize(fyne.NewSize(300, 400))

		// initialize the widgets
		titleEntry := widget.NewEntry()
		descEntry := widget.NewEntry()
		addBtn := widget.NewButton(
			"Add",
			func() {
				_ = todos.Append(models.LoadTodo(titleEntry.Text, descEntry.Text, false))
				database.Create(dbClient, titleEntry.Text, descEntry.Text, false)
				w.Close()
			},
		)

		// initialize the containers
		titleCtr := container.NewBorder(
			widget.NewLabel("Title:"),
			nil,
			nil,
			nil,
			titleEntry,
		)

		descCtr := container.NewBorder(
			widget.NewLabel("Description:"),
			nil,
			nil,
			nil,
			descEntry,
		)

		innerCtr := container.NewBorder(
			titleCtr,
			nil,
			nil,
			nil,
			descCtr,
		)

		// pass the content to the new window
		w.SetContent(
			container.NewBorder(
				nil,
				addBtn,
				nil,
				nil,
				innerCtr,
			),
		)

		w.SetOnClosed(
			func() {
				origin.Show()
			},
		)

		w.Show()
		origin.Hide()
	})
	return btn
}

// Initializes the button used for removing elements from the list.
func initializeDelBtn(newItemEntry *widget.Entry) *widget.Button {
	delBtn := widget.NewButton("Delete", func() {
		index := getTodoFromList(newItemEntry.Text)
		if index != -1 {
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
			"",
			b,
		))

		if err != nil {
			panic(err)
		}

		database.Update(dbClient, lbl.Text, b)
	})
}
