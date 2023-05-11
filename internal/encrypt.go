package internal

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
	"io"
)

func Encrypt(password string, in io.Reader, out io.Writer) error {
	salt, err := generateRandomString()
	if err != nil {
		return err
	}

	// Write the salt to the output
	_, err = out.Write([]byte(salt))
	if err != nil {
		return err
	}

	key := pbkdf2.Key([]byte(password), []byte(salt), iterations, keyLength, sha256.New)

	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	// Write the IV to the output
	_, err = out.Write(iv)
	if err != nil {
		return err
	}

	mode := cipher.NewCFBEncrypter(block, iv)

	//encoder := hex.NewEncoder(out)
	writer := &cipher.StreamWriter{S: mode, W: out}

	defer func() {
		//defer encoder.Close()

		defer writer.Close()
	}()

	if _, err := io.Copy(writer, in); err != nil {
		return err
	}

	return nil
}
