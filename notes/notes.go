package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func main() {
	// Read the JSON from the file
	noteJson, err := os.ReadFile("note.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create a new note to hold the decoded data
	var note Note

	// Decode the JSON to the note
	err = json.Unmarshal(noteJson, &note)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the note
	fmt.Printf("Title: %s\nContent: %s\nCreated At: %s\n", note.Title, note.Content, note.CreatedAt)
	fmt.Println("-------------------------------------------------")
}
