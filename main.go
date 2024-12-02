package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"massabatic.com/go-basics/notes"
	"massabatic.com/go-basics/todo"
)

type saver interface {
	Save() error
}

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
	err = saveData(userTodo)

	if err != nil {
		fmt.Print("Saving Todo Failed")
		return
	}
	fmt.Print("To do saved successfully")

	userNote.Display()
	err = saveData(userNote)

	if err != nil {
		fmt.Print("Saving Note Failed")
		return
	}
	fmt.Print("Note saved successfully")
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		return err
	}
	return nil
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
