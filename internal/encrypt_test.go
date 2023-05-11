package internal

import (
	"bytes"
	"crypto/aes"
	"strings"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	password := "my-password"
	plaintext := "Hello, World!"
	in := strings.NewReader(plaintext)
	encrypted := &bytes.Buffer{}

	err := Encrypt(password, in, encrypted)
	if err != nil {
		t.Fatalf("Error encrypting: %v", err)
	}

	decrypted := &bytes.Buffer{}
	err = Decrypt(password, bytes.NewReader(encrypted.Bytes()), decrypted)
	if err != nil {
		t.Fatalf("Error decrypting: %v", err)
	}

	decryptedText := decrypted.String()
	if decryptedText != plaintext {
		t.Errorf("Decrypted text does not match the original plaintext. Decrypted: %s, original: %s", decryptedText, plaintext)
	}
}

func TestEncrypt(t *testing.T) {
	password := "my-password"
	plaintext := "Hello, World!"
	in := strings.NewReader(plaintext)
	out := &bytes.Buffer{}

	err := Encrypt(password, in, out)
	if err != nil {
		t.Fatalf("Error encrypting: %v", err)
	}

	ciphertext := out.Bytes()

	if len(ciphertext) <= aes.BlockSize+16 { // Salt + IV + at least one block
		t.Errorf("Ciphertext too short")
	}

	if bytes.Contains(ciphertext[aes.BlockSize+16:], []byte(plaintext)) {
		t.Errorf("Plaintext found in ciphertext")
	}
}
