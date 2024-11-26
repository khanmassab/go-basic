package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title string
	Content string
	CreatedAt time.Time
}

func (note Note) Display() {
	fmt.Println("Title: ", note.Title)
	fmt.Println(note.Content)
	fmt.Println("---------------------")
}
 
func New (title, content string) (Note, error) {
	if title == "" || content== "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Save () error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
