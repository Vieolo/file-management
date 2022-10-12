package filemanagement_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	filemanagement "github.com/vieolo/file-management"
)

func TestDownload(t *testing.T) {
	f1, e1 := filemanagement.DownloadFileToBase64("https://images.unsplash.com/photo-1665493198207-49f755bc5c62?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=1112&q=80", false)
	assert.Equal(t, nil, e1)
	assert.Equal(t, filemanagement.ReadFileToBase64NoError("./files/image_1.jpeg"), f1)
}
