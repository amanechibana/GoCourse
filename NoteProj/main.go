package main

import (
	"fmt"
	"example.com/note/note"
)

func main() {
	title, content:= getNoteData()

	note.New(title,content)
	
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) (string) {
	fmt.Println(prompt)
	var value string
	fmt.Scan(&value)

	return value
}
