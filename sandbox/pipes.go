package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Create a pipe
	r, w, _ := os.Pipe()

	// Save the original stdout
	origStdout := os.Stdout

	// Redirect stdout to the write end of the pipe
	os.Stdout = w

	// Call a function that writes to stdout
	fmt.Println("Hello, world!")

	// Close the writer
	w.Close()

	// Restore stdout
	os.Stdout = origStdout

	// Read from the read end of the pipe
	out, _ := io.ReadAll(r)

	// Print the captured output
	fmt.Printf("Captured: %s", out)
}
