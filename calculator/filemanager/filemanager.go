package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, errors.New("error opening file: " + err.Error())
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		file.Close()
		return nil, errors.New("error reading file: " + err.Error())
	}

	file.Close()
	return lines, nil
}

// WriteJSON writes the data to a file at the specified path.
// If the file does not exist, it will be created.
// If the file already exists, it will be overwritten.
// If an error occurs, it will be returned.
func WriteJSON(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return errors.New("error creating file: " + err.Error())
	}

	jsonEncoder := json.NewEncoder(file)
	err = jsonEncoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("error encoding data: " + err.Error())
	}

	file.Close()
	return nil
}
