package notes

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title string
	content string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Println("Title: ", note.title)
	fmt.Println(note.content)
	fmt.Println("---------------------")
}
 
func New (title, content string) (Note, error) {
	if title == "" || content== "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		title: title,
		content: content,
		createdAt: time.Now(),
	}, nil
}
