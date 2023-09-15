package commands

import (
	"encoding/hex"
	"github.com/dbubel/passman/internal"
	"os"
)

type DecryptCmd struct{}

func (c *DecryptCmd) Help() string {
	return "Ex) cat somefile.txt | pman > somefile.encrypted"
}

func (c *DecryptCmd) Synopsis() string {
	return "decrypts a stream from os.Stdin"
}

func (c *DecryptCmd) Run(args []string) int {
	hexEncoder := hex.NewDecoder(os.Stdin)
	internal.Decrypt(os.Getenv("PW"), hexEncoder, os.Stdout)
	return 0
}
