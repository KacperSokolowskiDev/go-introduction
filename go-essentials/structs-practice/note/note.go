package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	// `` are used to add structs tags, metadata. This can be picked up by other packages (ex: JSON) to replace parameter name with defined tag
	Title string `json:"title"`
	Content string `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		Title: title,
		Content: content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the following content:\n %v\n\n", note.Title, note.Content)
}

func (note Note) Save() error {
	//replacing all blank spaces in title by underscores
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	//putting all characters in title to lower case and adding .json filetype
	fileName = strings.ToLower(fileName) + ".json"
	// this method can only access the available parameters so the value in the struct need to start with a capital letter
	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}