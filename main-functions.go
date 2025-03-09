package main

import (
	"bufio"
	"os"

	"todo.com/m/notes"
)

func HandleClearFlag(file *os.File) {
	todos := []notes.Todo{}

	jsonData := notes.ProcessTodosToJsonData(todos)
	notes.WriteTodosToFile(file, jsonData)
}

func HandleAddFlag(reader *bufio.Reader, file *os.File, title string) {
	var todo notes.Todo
	todo.CreateNewTodo(reader, title)

	todos := notes.GetTodosData(file)
	todo.GetNextId(todos)

	todos = append(todos, todo)

	jsonData := notes.ProcessTodosToJsonData(todos)
	notes.WriteTodosToFile(file, jsonData)
}

func HandleListFlag(file *os.File) {
	todos := notes.GetTodosData(file)
	notes.ProcessTodosToJsonData(todos)
}

func HandleDoneFlag(file *os.File, id int) {
	todos := notes.GetTodosData(file)

	for i, todo := range todos {
		if todo.Id == id {
			todo.MarkAsDone()
			todos[i] = todo
			break
		}
	}

	jsonData := notes.ProcessTodosToJsonData(todos)
	notes.WriteTodosToFile(file, jsonData)
}
