package filemanagement

import (
	"encoding/base64"
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
)

// Downloads the file from `fileURL` and writes to the disk
//
// The new file will be named as `diskFileName` and will be saved to `diskFilePath`
//
// If the permission of the file should be changed, it can be passes as `chmodPermission`, e.g. 0777
// The return values are -> (file name, file path, error)
func DownloadFileToDisk(diskFilePath string, diskFileName, fileURL string, chmodPermission *fs.FileMode) (string, string, error) {

	// Create the file
	out, err := os.Create(diskFilePath + diskFileName)
	if err != nil {
		return "", "", err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(fileURL)
	if err != nil {
		return "", "", err
	}
	defer resp.Body.Close()

	// Check server response
	if resp.StatusCode != http.StatusOK {
		return "", "", fmt.Errorf("bad status: %s", resp.Status)
	}

	// Writer the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", "", err
	}

	fPath := diskFilePath + diskFileName

	if chmodPermission != nil {
		chmodErr := os.Chmod(fPath, *chmodPermission)
		if chmodErr != nil {
			return "", "", chmodErr
		}
	}

	return diskFileName, diskFilePath + diskFileName, nil
}

// Downloads the file from `fileURL` and converts it into bytes
func DownloadFileToBytes(fileURL string) ([]byte, error) {
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

// Downloads the file from `fileURL` and converts it into base64
//
// The base64 can have a prefix, e.g. `data:image/jpeg;base64,`
//
// This prefix can be omitted or included using `includeDataPrefix`
func DownloadFileToBase64(fileURL string, includeDataPrefix bool) (string, error) {
	bytes, err := DownloadFileToBytes(fileURL)
	if err != nil {
		return "", err
	}

	var base64Encoding string
	mimeType := http.DetectContentType(bytes)
	base64Encoding += fmt.Sprintf("data:%v;base64,", mimeType)

	encoded := base64.StdEncoding.EncodeToString(bytes)

	if includeDataPrefix {
		return base64Encoding + encoded, nil
	}

	return encoded, nil
}

// Downloads a file from `fileURL` and converts it into a temporary `os.File`
func DownloadFileToFile(fileURL string) (*os.File, error) {
	resp, err := http.Get(fileURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP error: %d", resp.StatusCode)
	}

	// Create a temporary file to store the downloaded content
	tempFile, err := os.CreateTemp("", "downloaded_file")
	if err != nil {
		return nil, err
	}
	defer tempFile.Close()

	// Copy the response body to the temporary file
	_, err = io.Copy(tempFile, resp.Body)
	if err != nil {
		return nil, err
	}

	return tempFile, nil
}
