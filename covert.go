package filemanagement

import "encoding/base64"

// Converts byte array to Base 64
func ByteToBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
