package database

import (
	"encoding/json"
	"github.com/supabase-community/supabase-go"
	"todolist.com/models"
)

func InitializeClient() (*supabase.Client, error) {
	client, err := supabase.NewClient(API_URL, API_SERVICE_KEY, &supabase.ClientOptions{})

	if err != nil {
		return nil, err
	}

	return client, nil
}

func Read(client *supabase.Client) []models.Todo {
	data, _, err := client.From("todo").Select("*", "", false).Execute()

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
