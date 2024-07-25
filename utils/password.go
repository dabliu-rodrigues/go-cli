package utils

import (
	crytporand "crypto/rand"
	"encoding/base64"
	"fmt"
)

func GeneratePassword(size int) string {
	if size < 8 {
		fmt.Println("Size should be equal or greater than 8")
		return ""
	}

	b := make([]byte, size)
	if _, err := crytporand.Read(b); err != nil {
		fmt.Println("Error trying to generate random password")
		return ""
	}
	return base64.StdEncoding.EncodeToString(b)
}
