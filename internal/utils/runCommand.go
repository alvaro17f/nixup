package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func RunCommand(format string, args ...interface{}) (string, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "bash"
	}

	cmdStr := fmt.Sprintf(format, args...)
	cmd := exec.Command(shell, "-c", cmdStr)
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()

	return string(output), err
}
