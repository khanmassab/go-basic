package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"massabatic.com/go-basics/notes"
)

func main() {
	noteTitle, noteContent := getNoteData()

	userNote, err := notes.New(noteTitle, noteContent)
	if err != nil {
		fmt.Print(err)
		return
	}
	userNote.Display()
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
