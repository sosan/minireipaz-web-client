// generated by IA
package common

import (
	// "bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	// "fmt"
	"io"
)

func EncryptAESGCM(plaintext, key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("key must be 32 bytes for AES-256")
	}

	// Generate a random IV
	iv := make([]byte, 12)
	_, err := io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}

	// Create a new AES cipher in GCM mode
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}

	// Encrypt the plaintext
	ciphertext := gcm.Seal(nil, iv, plaintext, nil)

	// Combine IV and ciphertext
	combined := append(iv, ciphertext...)

	// Base64 encode the combined data
	encoded := base64.StdEncoding.EncodeToString(combined)

	return []byte(encoded), nil
}

func DecryptAESGCM(ciphertext, key []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("key must be 32 bytes for AES-256")
	}

	// Base64 decode the input
	combined, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return nil, err
	}

	// Extract IV and ciphertext
	if len(combined) < 12 {
		return nil, errors.New("ciphertext too short")
	}
	iv := combined[:12]
	ciphertextData := combined[12:]

	// Create a new AES cipher in GCM mode
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}

	// Decrypt the ciphertext
	plaintext, err := gcm.Open(nil, iv, ciphertextData, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
