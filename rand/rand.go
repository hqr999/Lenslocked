package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func Bytes(n int) ([]byte, error) {
	b := make([]byte, n)
	nRead, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("bytes: %w", err)
	}
	if nRead < n {
		return nil, fmt.Errorf("bytes:did not read enough random bytes")
	}

	return b, nil
}

// String retorna uma string randômica(pseudo) usando crypto/rand
// n é o número de bytes sendo usado para gerar um string aleatória
func String(n int) (string, error) {
	b, err := Bytes(n)
	if err != nil {
		return "", fmt.Errorf("string: %w", err)
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

const sessionTokenBytes = 32

func SessionToken() (string, error) {
	return String(sessionTokenBytes)
}
