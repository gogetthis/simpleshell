package utils

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("enter your command : ")
	scanner := bufio.NewScanner(os.Stdin)

	var response string
	for scanner.Scan() {
		response = scanner.Text()
		break
	}
	io.WriteString(w, response)
}
