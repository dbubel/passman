package commands

import (
	"bufio"
	"github.com/dbubel/passman/internal"
	"os"
)

type EncryptCmd struct {
}

func (c *EncryptCmd) Help() string {
	return ""
}

func (c *EncryptCmd) Synopsis() string {
	return "Runs the cohesion content API server"
}

func (c *EncryptCmd) Run(args []string) int {
	internal.Encrypt(os.Getenv("PW"), bufio.NewReader(os.Stdin), os.Stdout)
	return 0
}
