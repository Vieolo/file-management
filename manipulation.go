package filemanagement

import (
	"fmt"
	"os"
	"strings"
)

// Replaces a certain text in a file and writes it back to the disk
func ReplaceTextInFile(filePath string, oldText string, newText string) error {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(fileContent), "\n")

	for i, line := range lines {
		if strings.Contains(line, oldText) {
			lines[i] = strings.Replace(lines[i], oldText, newText, 1)
		}
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile(filePath, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}

// Inserts a text in a file after a landmark text
func InsertTextInFile(filePath string, toBeInserted string, landmark string, after bool) error {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(fileContent), "\n")

	for i, line := range lines {
		if strings.TrimSpace(line) == landmark {
			if after {
				lines[i] = fmt.Sprintf("%v\n%v", line, toBeInserted)
			} else {
				lines[i] = fmt.Sprintf("%v\n%v", toBeInserted, line)
			}
		}
	}
	output := strings.Join(lines, "\n")
	err = os.WriteFile(filePath, []byte(output), 0644)
	if err != nil {
		return err
	}

	return nil
}

// Adds new text to the bottom of the file and writes it back to the disk
func AppendToFile(filePath string, newContent string) error {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if _, err := f.Write([]byte(newContent)); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}

	return nil
}
