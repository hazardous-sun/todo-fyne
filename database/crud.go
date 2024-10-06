package database

import (
	"encoding/json"
	"github.com/supabase-community/supabase-go"
	"todolist.com/models"
)

func InitializeClient() *supabase.Client {
	client, err := supabase.NewClient(API_URL, API_SERVICE_KEY, &supabase.ClientOptions{})

	if err != nil {
		panic(err)
	}

	return client
}

func Read(client *supabase.Client, filters models.Filters) []models.Todo {
	flags := filters.GetFlags()
	var data []byte
	var err error
	switch flags[0] {
	case false:
		switch flags[1] {
		case false:
			data, _, err = client.From("todo").Select("*", "", false).Execute()
		case true:
			data, _, err = client.From("todo").Select("*", "", false).Eq("checked", "false").Execute()
		}
	case true:
		switch flags[1] {
		case false:
			data, _, err = client.From("todo").Select("*", "", false).Eq("checked", "true").Execute()
		case true:
			data, _, err = client.From("todo").Select("*", "", false).Execute()
		}
	}

	if err != nil {
		panic(err)
	}

	todos := todoArrFromByteArr(data)

	return todos
}

func todoArrFromByteArr(arr []byte) []models.Todo {
	var todos []models.Todo
	err := json.Unmarshal(arr, &todos)

	if err != nil {
		panic(err)
	}

	return todos
}

func Create(client *supabase.Client, text string, check bool) {
	todo := map[string]interface{}{
		"description": text,
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
		panic(err)
	}
}

func Update(client *supabase.Client, text string, check bool) {
	todo := map[string]interface{}{
		"description": text,
		"checked":     check,
	}
	_, _, err := client.From("todo").Update(
		todo,
		"TEST UPDATE",
		"0",
	).Eq("description", text).Execute()

	if err != nil {
		panic(err)
	}
}

func Delete(client *supabase.Client, text string) {

	_, _, err := client.From("todo").Delete(
		"TEST DELETE",
		"exact",
	).Eq("description", text).Execute()

	if err != nil {
		panic(err)
	}
}
