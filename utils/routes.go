package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("enter your command (%v) : ", r.Header.Get("User-Agent"))
	scanner := bufio.NewScanner(os.Stdin)

	var response string
	for scanner.Scan() {
		response = scanner.Text()
		break
	}
	io.WriteString(w, response)
}

func HandleResponse(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		io.WriteString(w, "Update available")
	}

	content, _ := io.ReadAll(r.Body)
	fmt.Println(string(content))

	io.WriteString(w, "Already upto date")
}
