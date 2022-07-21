package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Book struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

func main() {
	book := Book{
		Title:       "Learning Go",
		Description: "Learning about Go",
	}
	data, err := json.Marshal(book)
	if err != nil {
		panic(err)
	}

	f, err := os.CreateTemp("", "books-*.json")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Writing to file %s\n", f.Name())
	os.WriteFile(f.Name(), data, 0644)
}
