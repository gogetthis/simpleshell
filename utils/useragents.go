package utils

import (
	"bufio"
	"math/rand"
	"os"
)

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
