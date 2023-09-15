package commands

import (
	"bufio"
	"encoding/hex"
	"github.com/dbubel/passman/internal"
	"os"
)

type EncryptCmd struct{}

func (c *EncryptCmd) Help() string {
	return "Ex) cat somefile.encrypted | pman > somefile.txt"
}

func (c *EncryptCmd) Synopsis() string {
	return "encrypts a stream from os.Stdin"
}

func (c *EncryptCmd) Run(args []string) int {
	hexEncoder := hex.NewEncoder(os.Stdout)
	internal.Encrypt(os.Getenv("PW"), bufio.NewReader(os.Stdin), hexEncoder)
	return 0
}
