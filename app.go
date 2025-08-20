package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	note "example.com/notes"
)

func main() {
	// This is a placeholder for the main function.
	// You can add your code here.
	title := getInput("Enter the title of the note: ")
	content := getInput("Enter the content of the note: ")

	newNote := note.New(title, content)
	// fmt.Print(notes.Content)
	newNote.Display()

	err := newNote.Save()
	if err != nil {
		fmt.Println("Error saving note:", err)
	} else {
		fmt.Println("Note saved successfully!")
	}
}

func getInput(display string) string {
	var input string
	fmt.Print(display)
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}
