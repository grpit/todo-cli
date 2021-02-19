package utils

import (
	"encoding/json"
	"os"
)

func WriteJSONToFile(project interface{}, path string) error {
	file, _ := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")

	if err := encoder.Encode(project); err != nil {
		return err
	}

	return nil
}
