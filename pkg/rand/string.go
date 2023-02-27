package rand

import (
	"crypto/rand"
	"encoding/base64"
)

// String from cypto/rand
func String(len int) (string, error) {
	b := make([]byte, len)
	if _, err := rand.Read(b); err != nil {
		return "", err

	}
	return base64.URLEncoding.EncodeToString(b), nil
}
