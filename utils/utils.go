package utils

import (
	"os"
	"os/exec"
)

func Run(command string) error {
	cmd := exec.Command(`bash`, `-c`, command)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
