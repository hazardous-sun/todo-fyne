package database

import (
	"encoding/json"
	"fmt"
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

func Read(client *supabase.Client) []models.Todo {
	data, count, err := client.From("todo").Select("*", "", false).Execute()

	fmt.Println(count)

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
		"TEST",
		"0",
	).Execute()

	if err != nil {
		panic(err)
	}
}
