package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	// Variable decleration
	var noteTitle string
	var noteContent string
	var userNote *note.Note

	getUserInput("Note title:", &noteTitle)
	getUserInput("Note content:", &noteContent)

	userNote, err := note.New(noteTitle, noteContent)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed.")
		fmt.Println(err)
		return
	}

	fmt.Println("Saving succeeded!")
}

func getUserInput(prompt string, value *string) {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		*value = ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	*value = text
}
