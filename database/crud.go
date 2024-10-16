// Package database :
// Contains the CRUD functions for the project.
package database

import (
	"encoding/json"
	"fmt"
	"github.com/supabase-community/supabase-go"
	"todolist.com/models"
)

// InitializeClient :
// Returns an instance to the client of Supabase, that is used to run the CRUD functions.
// It uses an API_URL and an API_SERVICE_KEY, that should be maintained inside the project.
func InitializeClient() *supabase.Client {
	client, _ := supabase.NewClient(API_URL, API_SERVICE_KEY, &supabase.ClientOptions{})
	return client
}

// Read :
// Returns an array of models.TodoItem with a selection of values from the database, based on the values of models.Filters.
// It will select "*" when all values of models.Filters are true AND when all of the values are false.
func Read(client *supabase.Client, filters models.Filters, table string) ([]models.TodoItem, error) {
	flags := filters.GetFlags()
	var data []byte
	var err error
	switch flags[0] {
	case false:
		switch flags[1] {
		case false:
			data, _, err = client.From(table).Select("*", "", false).Execute()
		case true:
			data, _, err = client.From(table).Select("*", "", false).Eq("checked", "false").Execute()
		}
	case true:
		switch flags[1] {
		case false:
			data, _, err = client.From(table).Select("*", "", false).Eq("checked", "true").Execute()
		case true:
			data, _, err = client.From(table).Select("*", "", false).Execute()
		}
	}

	if err != nil {
		return nil, err
	}

	todos := todoArrFromByteArr(data)

	return todos, nil
}

// Parses the values received from the database into an array of models.TodoItem.
func todoArrFromByteArr(arr []byte) []models.TodoItem {
	var todos []models.TodoItem
	err := json.Unmarshal(arr, &todos)

	if err != nil {
		panic(err)
	}

	return todos
}

// Create :
// Inserts a new value into the database.
func Create(client *supabase.Client, title string, description string, check bool) error {
	todo := map[string]interface{}{
		"title":       title,
		"description": description,
		"checked":     check,
	}
	_, _, err := client.From("todo").Insert(
		todo,
		false,
		"ERROR",
		"TEST CREATE",
		"0",
	).Execute()

	if err != nil {
		fmt.Println("error during insertion of value'", todo["title"], todo["description"], todo["checked"], "' to the database:", err)
		return err
	}

	return nil
}

// Update :
// Updates the value of "checked" inside the database for the row that matches the description of the todo item.
func Update(client *supabase.Client, title string, description string, check bool) {
	todo := map[string]interface{}{
		"title":       title,
		"description": description,
		"checked":     check,
	}
	_, _, err := client.From("todo").Update(
		todo,
		"TEST UPDATE",
		"0",
	).Eq("title", title).Execute()

	if err != nil {
		panic(err)
	}
}

// Delete :
// Removes from the database the row that matches the description provided.
func Delete(client *supabase.Client, title string) {

	_, _, err := client.From("todo").Delete(
		"TEST DELETE",
		"exact",
	).Eq("title", title).Execute()

	if err != nil {
		panic(err)
	}
}
