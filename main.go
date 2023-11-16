package main

import (
	"bufio"
	"strings"

	"fmt"

	"os"

	"main/dictionary"
)

func main() {
	d := dictionary.New()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command (add/list/exit): ")
		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
		case "add":
			actionAdd(d, reader)
		case "list":
			actionList(d)
		case "exit":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid command. Please try again.")
		}
	}
}

func actionAdd(d *dictionary.Dictionary, reader *bufio.Reader) {

	fmt.Print("Enter word: ")

	word, _ := reader.ReadString('\n')

	fmt.Print("Enter definition: ")

	definition, _ := reader.ReadString('\n')

	d.Add(word, definition)

	fmt.Println("Word added successfully.")

}

func actionDefine(d *dictionary.Dictionary, reader *bufio.Reader) {

}

func actionRemove(d *dictionary.Dictionary, reader *bufio.Reader) {

}

func actionList(d *dictionary.Dictionary) {

	wordList, entries := d.List()

	if len(wordList) == 0 {
		fmt.Println("No words in the dictionary.")
		return
	}

	fmt.Println("Word List:", strings.Join(wordList, ", "))

	fmt.Println("Entries:")
	for _, word := range wordList {
		entry, _ := entries[word]
		fmt.Printf("%s: %s\n", word, entry)

	}
}
