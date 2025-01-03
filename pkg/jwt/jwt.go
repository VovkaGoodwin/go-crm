package jwt

import (
	"crypto/rand"
	"fmt"
)

func GenerateTokenId(length int32) (id string, err error) {
	buf := make([]byte, length)
	_, err = rand.Read(buf)
	return fmt.Sprintf("%x", buf), err
}
