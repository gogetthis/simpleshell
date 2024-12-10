package utils

import (
	"os/exec"
	"runtime"
)

func (j *Job) Execute() bool {

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd.exe", "/c", j.Command)
	case "darwin", "linux":
		cmd = exec.Command("/bin/bash", "-c", j.Command)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return false
	}

	j.Output = string(output)
	return true
}
