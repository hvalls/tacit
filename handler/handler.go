package handler

import (
	"bytes"
	"fmt"
	"os/exec"
)

const (
	BASH          = "/bin/bash"
	DEFAULT_SHELL = BASH
)

func Handle(shell string, scriptPath string, args []string) (string, string, error) {
	cmd := exec.Command(shell, append([]string{scriptPath}, args...)...)
	var stdoutBuffer, stderrBuffer bytes.Buffer
	cmd.Stdout = &stdoutBuffer
	cmd.Stderr = &stderrBuffer
	if err := cmd.Start(); err != nil {
		fmt.Println("Error 1", err)
		return "", "", err
	}
	if err := cmd.Wait(); err != nil {
		fmt.Println("Error 2", err)
		return "", "", err
	}
	return stdoutBuffer.String(), stderrBuffer.String(), nil
}
