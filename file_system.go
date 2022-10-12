package filemanagement

import "os"

// Reads a file from the disk and converts to the base64
// This function does not return any error. if any errors, it returns an empty string
func ReadFileToBase64NoError(filePath string) string {
	f, e := os.ReadFile(filePath)

	if e != nil {
		return ""
	}

	return ByteToBase64(f)
}
