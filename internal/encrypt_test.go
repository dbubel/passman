package internal

import (
	"bytes"
	"testing"
)

// Testing generateRandomString function
func TestGenerateRandomString(t *testing.T) {
	s, err := generateRandomString()
	if err != nil {
		t.Fatal(err)
	}

	if len(s) != 16 {
		t.Errorf("Expected string of length 16, got %d", len(s))
	}
}

// Testing Encrypt function
func TestEncrypt(t *testing.T) {
	password := "password"
	in := bytes.NewBufferString("This is a test message")
	out := &bytes.Buffer{}

	err := Encrypt(password, in, out)
	if err != nil {
		t.Fatal(err)
	}

	if out.Len() == 0 {
		t.Errorf("Expected output to be non-empty")
	}
}

// Testing Decrypt function
func TestDecrypt(t *testing.T) {
	password := "password"
	message := "This is a test message"

	in := bytes.NewBufferString(message)
	encrypted := &bytes.Buffer{}
	err := Encrypt(password, in, encrypted)
	if err != nil {
		t.Fatal(err)
	}

	decrypted := &bytes.Buffer{}
	err = Decrypt(password, bytes.NewBuffer(encrypted.Bytes()), decrypted)
	if err != nil {
		t.Fatal(err)
	}

	if decrypted.String() != message {
		t.Errorf("Expected %s, got %s", message, decrypted.String())
	}
}
