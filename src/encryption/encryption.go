package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

const key = "01234567898123456789012345678901"

var (
	ErrCiphertextShort = errors.New("ciphertext too short")
)

func Encrypt(plaintext []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gen, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gen.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gen.Seal(nonce, nonce, plaintext, nil), nil
}

func Decrypt(ciphertext []byte) ([]byte, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}

	gen, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gen.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, ErrCiphertextShort
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gen.Open(nil, nonce, ciphertext, nil)
}

func ToBase64(plaintext []byte) string {
	return base64.RawStdEncoding.EncodeToString(plaintext)
}

func FromBase64(ciphertext string) ([]byte, error) {
	return base64.RawStdEncoding.DecodeString(ciphertext)
}
