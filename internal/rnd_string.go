package internal

import (
	"crypto/rand"
)

//const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//
//func generateRandomString() (string, error) {
//	length, err := rand.Int(rand.Reader, big.NewInt(11)) // Get a random length between 0 and 10
//	if err != nil {
//		return "", err
//	}
//	length.Add(length, big.NewInt(10)) // Add 10 to get a length between 10 and 20
//
//	b := make([]byte, length.Int64())
//	for i := range b {
//		var index *big.Int
//		index, err = rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
//		if err != nil {
//			return "", err
//		}
//		b[i] = charset[index.Int64()]
//	}
//
//	return string(b), nil
//}
