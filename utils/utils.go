package utils

import "encoding/base64"

//Base64Str fwefdw
func Base64Str(msg string) string {
	return base64.StdEncoding.EncodeToString([]byte(msg))
}
