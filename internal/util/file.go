package util

import (
	"errors"
	"fmt"
	"os"
)

func WriteToFile(filename, data string) error {
	var file *os.File
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			file, _ = os.Create(filename)
		}
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

func AppendToFile(filename, data string) error {
	var file *os.File
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			file, _ = os.Create(filename)
		}
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}