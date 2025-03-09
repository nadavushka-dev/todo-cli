package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"todo.com/m/notes"
)

func main() {
	add := flag.String("a", "", "Add a new todo")
	ls := flag.Bool("ls", false, "List all todos")
	done := flag.Int("d", 0, "Mark selected task as done")
	clear := flag.Bool("c", false, "Clear all todos")

	flag.Parse()

	reader := bufio.NewReader(os.Stdin)
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

	fmt.Println("Usage: ")
	flag.PrintDefaults()

}
