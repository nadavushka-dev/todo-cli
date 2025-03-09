package notes

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadTodosFile() *os.File {
	file, err := os.OpenFile("todos.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("unable to open file. \n", err)
	}

	return file
}

func GetTodosData(file *os.File) []Todo {
	var todos []Todo
	data, _ := io.ReadAll(file)

	if len(data) > 0 {
		json.Unmarshal(data, &todos)
	}

	return todos
}

func WriteTodosToFile(file *os.File, jsonData []byte) {
	file.Truncate(0)
	_, err := file.WriteAt(jsonData, 0)
	if err != nil {
		fmt.Println("problems writing to the file")
	}
}

func ProcessTodosToJsonData(todos []Todo) []byte {
	jsonData, _ := json.MarshalIndent(todos, "", " ")
	fmt.Printf("All todos: \n%s\n", string(jsonData))

	return jsonData
}
