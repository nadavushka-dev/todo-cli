package main

import (
	"bufio"
	"fmt"
	"os"

	"todo.com/m/notes"
)

func HandleClearFlag(file *os.File) {
	todos := []notes.Todo{}

	jsonData := notes.ProcessTodosToJsonData(todos)
	notes.WriteTodosToFile(file, jsonData)
	fmt.Println()
}

func HandleAddFlag(reader *bufio.Reader, file *os.File, title string) {
	var todo notes.Todo
	todo.CreateNewTodo(reader, title)

	todos := notes.GetTodosData(file)
	todo.GetNextId(todos)

	todos = append(todos, todo)

	jsonData := notes.ProcessTodosToJsonData(todos)
	notes.WriteTodosToFile(file, jsonData)
	fmt.Println()
}

func HandleListFlag(file *os.File) {
	todos := notes.GetTodosData(file)
	notes.ProcessTodosToJsonData(todos)
	fmt.Println()
}

func HandleDoneFlag(file *os.File, id int) {
	todos := notes.GetTodosData(file)
	fmt.Println("todos: ", todos)

	for i, todo := range todos {
		fmt.Println("todo: ", todo, "id: ", id)
		if todo.Id == id {
			fmt.Println("here")
			todo.MarkAsDone()
			todos[i] = todo
			break
		}
	}

	jsonData := notes.ProcessTodosToJsonData(todos)
	notes.WriteTodosToFile(file, jsonData)
	fmt.Println()
}
