package note

import (
	"errors"
	"time"
)

type Note struct {
	title string
	content string
	createdAt time.Time
}

func New (title, content string) (Note, error) {
	if title == "" || content== "" {
		return Note{}, errors.New("invalid input")
	}
	return Note{
		title: title,
		content: title,
		createdAt: time.Now(),
	}, nil
}

func main(){

}