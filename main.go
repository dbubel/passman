package main

import (
	"fmt"
	"github.com/dbubel/passman/commands"
	"github.com/mitchellh/cli"
	"os"
)

func main() {

	c := cli.NewCLI("passman", "v1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"encrypt": func() (cli.Command, error) {
			return &commands.EncryptCmd{}, nil
		},
	}

	_, err := c.Run()
	if err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Printf("Ciphertext: %s\n", hex.EncodeToString(ciphertext))
	//
	//// Decrypt the message
	//plaintext, err := internal.Decrypt(ciphertext, passwordStr)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Printf("Plaintext: %s\n", plaintext)
}
