package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"todo.com/m/notes"
)

func main() {
	add := flag.String("a", "", "Add a new todo")
	ls := flag.Bool("ls", false, "List all todos")
	done := flag.Int("d", 0, "Mark selected task as done")
	clear := flag.Bool("c", false, "Clear all todos")
	help := flag.Bool("h", false, "help command")

	reader := bufio.NewReader(os.Stdin)

	flag.Parse()

	if flag.NFlag() == 0 {
		RunNoFlags(reader)
		return
	}

	file := notes.ReadTodosFile()
	defer file.Close()

	if *clear {
		HandleClearFlag(file)
		return
	}

	if *done > 0 {
		HandleDoneFlag(file, *done)
		return
	}

	if *add != "" {
		if flag.NArg() > 0 {
			*add += " " + strings.Join(flag.Args(), " ")
		}
		HandleAddFlag(reader, file, *add)
		return
	}

	if *ls {
		HandleListFlag(file)
		return
	}

	if *help {
		fmt.Println("Usage: ")
		flag.PrintDefaults()
	} else {
	}
}

func RunNoFlags(reader *bufio.Reader) {
	var run bool = true
	for run {
		file := notes.ReadTodosFile()

		fmt.Println("select option: ")
		fmt.Println("1: list all todos")
		fmt.Println("2: add a todo")
		fmt.Println("3: complete todo")
		fmt.Println("4: clear todo list")
		fmt.Println("q: quite program")
		opt, _ := reader.ReadString('\n')
		opt = strings.TrimSpace(opt)

		if opt == "1" {
			HandleListFlag(file)
		} else if opt == "2" {
			fmt.Println("provide a title")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			HandleAddFlag(reader, file, title)
		} else if opt == "3" {
			fmt.Println("provide todo's id")
			id, _ := reader.ReadString('\n')
			nId, _ := strconv.Atoi(strings.TrimSpace(id))
			HandleDoneFlag(file, nId)
		} else if opt == "4" {
			HandleClearFlag(file)
		} else if opt == "q" {
			run = false
		}

		file.Close()
	}
}
