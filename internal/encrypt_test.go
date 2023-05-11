package internal

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
)

func TestEncryptDecrypt(t *testing.T) {
	password := "testpassword"
	inputData := "testdata"

	// Setup an input reader with the input data
	in := strings.NewReader(inputData)

	// Setup a buffer to capture the encrypted output
	encryptedOut := &bytes.Buffer{}

	err := Encrypt(password, in, encryptedOut)
	if err != nil {
		t.Fatalf("Expected no error during encryption, got: %v", err)
	}

	//t.Logf("%02x\n", encryptedOut)

	// Setup a reader with the encrypted data
	encryptedIn := bytes.NewReader(encryptedOut.Bytes())

	// Setup a buffer to capture the decrypted output
	decryptedOut := &bytes.Buffer{}

	err = Decrypt(password, encryptedIn, decryptedOut)
	if err != nil {
		t.Fatalf("Expected no error during decryption, got: %v", err)
	}

	decryptedData := decryptedOut.String()

	// Check if the decrypted data matches the original input
	if inputData != decryptedData {
		t.Fatalf("Expected decrypted data to be %v, got: %v", inputData, decryptedData)
	}
}

func TestEncrypt(t *testing.T) {
	password := "testpassword"
	inputData := "testdata"

	// Setup an input reader with the input data
	in := strings.NewReader(inputData)

	// Setup a buffer to capture the output
	out := &bytes.Buffer{}

	err := Encrypt(password, in, out)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	encryptedData := out.String()
	fmt.Fprintf(os.Stdout, "hex: %02x\n", encryptedData)

	//// Check if the output data is not empty and is base64 encoded
	//if len(encryptedData) == 0 {
	//	t.Fatalf("Expected output to not be empty")
	//}
	//
	//// The output should have at least the size of the salt and the IV, plus the base64 encoding of the input data.
	//// Note: this check assumes that the salt is of a fixed size.
	//expectedMinSize := len("randomSalt") + aes.BlockSize + base64.StdEncoding.EncodedLen(len(inputData))
	//if len(encryptedData) < expectedMinSize {
	//	t.Fatalf("Expected output size to be at least %v, got: %v", expectedMinSize, len(encryptedData))
	//}
}

const (
	testDataSize = 1000 * 1024 * 1024 // 100 MB
)

func BenchmarkEncrypt(b *testing.B) {
	// Create some test data
	testData := make([]byte, testDataSize)

	password := "password"

	// Create the input and output buffers
	in := bytes.NewReader(testData)
	out := &bytes.Buffer{}

	start := time.Now()

	for i := 0; i < b.N; i++ {
		// Reset the input reader and the output buffer
		in.Seek(0, 0)
		out.Reset()

		err := Encrypt(password, in, out)
		if err != nil {
			b.Fatal(err)
		}
	}

	elapsed := time.Since(start)
	throughput := float64(testDataSize*b.N) / elapsed.Seconds() / 1024 / 1024 // in MB/s
	b.Logf("Throughput: %.2f MB/s", throughput)
}
