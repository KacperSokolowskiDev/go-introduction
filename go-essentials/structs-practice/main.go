package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
  "example.com/note/todo"
)

//defines that a certain method exists, only name, parameters & return value, not defining the function body. Works like a contract for structs.
type saver interface {
  Save() error
}

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
  err = saveData(userNote)

  if err != nil {
    return
  }

  todoText := getUserInput("Todo text: ")

  todo, err := todo.New(todoText)

  if err != nil {
    fmt.Println(err)
    return
  }

  todo.Display()
  err = saveData(todo)

  if err != nil {
    return
  }
  
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func saveData(data saver) error {
  err := data.Save()

  if err != nil {
    fmt.Println("Saving failed")
    return err
  }

  fmt.Println("Saving succeeded")
  return nil
}

func getUserInput(prompt string) (string) {
	fmt.Print(prompt)
	//reader that takes data from command line
	reader := bufio.NewReader(os.Stdin)
	//defining at what byte/character it should stop reading. Here use single quotes for \n
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	//if no error, cleaning the text value from the special characters
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}