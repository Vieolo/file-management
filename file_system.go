package filemanagement

import (
	"os"
	"path/filepath"
	"strings"
)

// Reads a file from the disk and converts to the base64
// This function does not return any error. if any errors, it returns an empty string
func ReadFileToBase64NoError(filePath string) string {
	f, e := os.ReadFile(filePath)

	if e != nil {
		return ""
	}

	return ByteToBase64(f)
}

func GetDirFilesFlat(scrDir string) ([]string, error) {
	paths := []string{}

	entries, err := os.ReadDir(scrDir)
	if err != nil {
		return []string{}, err
	}
	for _, entry := range entries {
		currentPath := filepath.Join(scrDir, entry.Name())
		if !entry.IsDir() {
			paths = append(paths, currentPath)
		} else {
			subPaths, subErr := GetDirFilesFlat(currentPath)
			if subErr != nil {
				return []string{}, subErr
			}
			paths = append(paths, subPaths...)
		}

	}
	return paths, nil
}

// Finds the first file in a folder (and its children) that contains a text
func FindFirstFileInFolderContainingText(folderPath string, subStr string) (string, error) {
	targetFile := ""

	// The files in the target folder is searched for the target string
	folderFiles, ffErr := GetDirFilesFlat(folderPath)
	if ffErr != nil {
		return "", ffErr
	}

	for _, f := range folderFiles {
		fileContent, fcErr := os.ReadFile(f)
		if fcErr != nil {
			return "", fcErr
		}
		if strings.Contains(string(fileContent), subStr) {
			// The sub string is found in the file
			targetFile = f
			break
		}
	}

	if targetFile == "" {
		// The sub string is not found in any of the files in the target folder
		return "", &NotFoundError{Target: subStr}
	}

	return targetFile, nil
}
