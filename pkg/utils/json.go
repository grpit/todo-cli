package utils

import (
	"encoding/json"
	"errors"
	"os"
)

// CreateFileWithPerm creates a file with write permissions
func CreateFileWithPerm(path string) error {
	_, err := os.OpenFile(path, os.O_RDONLY, os.ModePerm)

	// return error if project file exists
	if err == nil {
		return errors.New("")
	}

	os.Create(path)
	return nil
}

// WriteJSONToFile writes any type to json in the .todo.json file
func WriteJSONToFile(project interface{}, path string) error {
	file, _ := os.OpenFile(path, os.O_WRONLY, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")

	if err := encoder.Encode(project); err != nil {
		return err
	}

	return nil
}
