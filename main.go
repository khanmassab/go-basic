package main

import (
	"errors"
	"fmt"

	"massabatic.com/go-basics/note"
)

func main() {
	noteTitle, noteContent := getNoteData()

	var note note
	userNote, err := note.New(noteTitle, noteContent)
	if err != nil {
		fmt.Print(err)
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

	var value string
	fmt.Scanln(&value)
	
	return value, 
}
