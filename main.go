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
		"decrypt": func() (cli.Command, error) {
			return &commands.DecryptCmd{}, nil
		},
	}

	_, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}
