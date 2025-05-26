package utils

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

func SaveToFile(filename string, data interface{}) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(filename, bytes, 0644)
}

func LoadFromFile(filename string, data interface{}) error {
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		empty := []byte("[]")
		err := os.WriteFile(filename, empty, 0644)
		if err != nil {
			return err
		}
	}

	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	return json.Unmarshal(bytes, data)
}