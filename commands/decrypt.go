package commands

import (
	"bufio"
	"encoding/hex"
	"github.com/dbubel/passman/internal"
	"os"
)

type DecryptCmd struct {
}

func (c *DecryptCmd) Help() string {
	return ""
}

func (c *DecryptCmd) Synopsis() string {
	return "Runs the cohesion content API server"
}

func (c *DecryptCmd) Run(args []string) int {
	hexEncoder := hex.NewEncoder(os.Stdout)
	internal.Encrypt(os.Getenv("PW"), bufio.NewReader(os.Stdin), hexEncoder)
	return 0
}
