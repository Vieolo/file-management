package filemanagement

import "encoding/base64"

func ByteToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
