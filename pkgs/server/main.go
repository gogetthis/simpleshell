package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"

	"github.com/gogetthis/simpleshell/utils"
)

func doNothing(w http.ResponseWriter, r *http.Request) {}

func main() {

	certPtr := flag.String("cert", "certs/server.crt", "enter the path to server.crt file")
	keyPtr := flag.String("key", "certs/server.key", "enter the path to server.key file")
	addrPtr := flag.String("addr", ":8443", "enter the port to start the server")

	flag.Parse()

	cert, _ := tls.LoadX509KeyPair(*certPtr, *keyPtr)

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	router := http.NewServeMux()

	router.HandleFunc("/favicon.ico", doNothing) // to avoid calling the handler twice
	router.HandleFunc("/", utils.HandleRoot)

	server := &http.Server{
		Addr:      *addrPtr,
		Handler:   router,
		TLSConfig: config,
	}
	fmt.Printf("Starting the server at %v\n\n", *addrPtr)

	err := server.ListenAndServeTLS("", "")
	if err != nil {
		panic(err)
	}
}
