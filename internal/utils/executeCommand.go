package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteCommand(format string, args ...interface{}) (string, error) {
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "bash"
	}

	cmdStr := fmt.Sprintf(format, args...)
	cmd := exec.Command(shell, "-c", cmdStr)
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()

	if err != nil {
		return "", fmt.Errorf("error executing command: %w", err)
	}

	return string(output), nil
}
