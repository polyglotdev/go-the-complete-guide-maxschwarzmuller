package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"example.com/structs/user"
)

func TestUserCreation(t *testing.T) {
	// Prepare the input
	input := strings.NewReader("Elijah\nHallan\n10/17/1999\n")

	// Create a pair of connected file descriptors
	r, w, _ := os.Pipe()

	// Replace os.Stdin temporarily
	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()
	os.Stdin = r

	// Copy the input to the write end of the pipe in a separate goroutine
	// to avoid blocking
	go func() {
		defer w.Close()
		_, err := io.Copy(w, input)
		if err != nil {
			t.Errorf("Error copying input to the write end of the pipe: %v", err)
		}
		if err := w.Close(); err != nil {
			t.Errorf("Error closing the write end of the pipe: %v", err)
		}
	}()

	// Call the function that reads from os.Stdin
	user, err := user.New()
	if err != nil {
		t.Errorf("createUser() error = %v", err)
		return
	}

	// Check the result
	if user.FirstName != "Elijah" || user.LastName != "Hallan" {
		t.Errorf("createUser() = %v, want %v", user, "Elijah Hallan")
	}
}
