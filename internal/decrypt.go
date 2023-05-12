package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
	"io"
)

func Decrypt(password string, in io.Reader, out io.Writer) error {
	salt := make([]byte, 16)
	_, err := io.ReadFull(in, salt)
	if err != nil {
		return err
	}

	key := pbkdf2.Key([]byte(password), salt, iterations, keyLength, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(in, iv); err != nil {
		return err
	}

	mode := cipher.NewCFBDecrypter(block, iv)

	writer := &cipher.StreamWriter{S: mode, W: out}

	defer writer.Close()

	_, err = io.Copy(writer, in)

	return err
}
