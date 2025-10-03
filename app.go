package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/functions"
	//note "example.com/notes"
)

func main() {
	functions.Mani()
	//arrays.ArraySlice()
	// This is a placeholder for the main function.
	// You can add your code here.
	// title := getInput("Enter the title of the note: ")
	// content := getInput("Enter the content of the note: ")

	// newNote := note.New(title, content)
	// // fmt.Print(notes.Content)
	// newNote.Display()

	// err := newNote.Save()
	// if err != nil {
	// 	fmt.Println("Error saving note:", err)
	// } else {
	// 	fmt.Println("Note saved successfully!")
	// }

	// values := [19]float64{1, 2, 3, 4, 5, 6, 7}
	// var productName [19]string
	// productName = [19]string{"Apple", "Banana", "Orange"}
	// fmt.Print(values)
	// fmt.Print(productName)
	// fmt.Print(len(values), cap(values))
	// fmt.Println(productName)
	// fmt.Print(productName[0:2])
	// selected := values[1:5]
	// superSelected := selected[1:3]
	// fmt.Print(selected)
	// fmt.Print(superSelected)
	// fmt.Print(len(superSelected), cap(superSelected))
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
