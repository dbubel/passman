package commands

import (
	"encoding/hex"
	"github.com/dbubel/passman/internal"
	"os"
)

type DecryptCmd struct{}

func (c *DecryptCmd) Help() string {
	return ""
}

func (c *DecryptCmd) Synopsis() string {
	return "Runs the cohesion content API server"
}

func (c *DecryptCmd) Run(args []string) int {
	hexEncoder := hex.NewDecoder(os.Stdin)
	internal.Decrypt(os.Getenv("PW"), hexEncoder, os.Stdout)
	return 0
}
