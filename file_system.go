package filemanagement

import "os"

func ReadFileToBase64NoError(filePath string) string {
	f, e := os.ReadFile(filePath)

	if e != nil {
		return ""
	}

	return ByteToBase64(f)
}
