package notes

import (
	"bufio"
	"strings"
)

type Status string

const (
	Open   Status = "open"
	Closed Status = "closed"
)

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status Status `json:"status"`
}

func (todo *Todo) CreateNewTodo(reader *bufio.Reader, title string) {
	todo.Title = strings.TrimSpace(title)
	todo.Status = Open
}

func (todo *Todo) GetNextId(todos []Todo) {
	var id int
	if len(todos) > 0 {
		id = todos[len(todos)-1].Id + 1
	} else {
		id = 1
	}

	todo.Id = id
}

func (todo *Todo) MarkAsDone() {
	todo.Status = Closed
}
