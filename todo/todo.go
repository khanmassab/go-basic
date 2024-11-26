package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const fileName = "todo.json"

type Todo struct {
	Text string `json:"text"` //struct tags "metadata processed by json standard library"
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}
 
func New (content string) (Todo, error) {
	if content== "" {
		return Todo{}, errors.New("invalid input")
	}
	return Todo{
		Text: content,
	}, nil
}

func (todo Todo) Save () error {
	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}
	
	return os.WriteFile(fileName, json, 0644)
}
