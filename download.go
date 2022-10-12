package filemanagement

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

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

func DownloadFileToBase64(fileURL string, includeDataPrefix bool) (string, error) {
	bytes, err := DownloadFileToBytes(fileURL)
	if err != nil {
		return "", err
	}

	var base64Encoding string
	mimeType := http.DetectContentType(bytes)
	base64Encoding += fmt.Sprintf("data:%v;base64", mimeType)

	encoded := base64.StdEncoding.EncodeToString(bytes)

	if includeDataPrefix {
		return base64Encoding + encoded, nil
	}

	return encoded, nil
}
