package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"massabatic.com/go-basics/notes"
	"massabatic.com/go-basics/todo"
)

func main() {
	noteTitle, noteContent := getNoteData()

	noteText := getUserInput("Input ToDo Text: ")

	userNote, err := notes.New(noteTitle, noteContent)

	if err != nil {
		fmt.Print(err)
		return
	}

	userTodo, err := todo.New(noteText)

	if err != nil {
		fmt.Print(err)
		return
	}

	userTodo.Display()
	err = userTodo.Save()
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println(err)
		return
	}
}

func getNoteData () (string, string) {
	noteTitle := getUserInput("Input the note title: ")
	noteContent := getUserInput("Input the note content: ")

	return noteTitle, noteContent
}

func getUserInput(str string) string {
	fmt.Print(str)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	//for windows line break = \r\n
	text = strings.TrimSuffix(text, "\r")
	
	return text
}
