package main

import (
	"balun_homework_1/foundation/network"
	"bufio"
	"flag"
	"fmt"
	"os"
)

var address string

func init() {
	flag.StringVar(&address, "addr", "127.0.0.1:3223", "database server address")
}

func main() {
	flag.Parse()

	client, err := network.NewClient(address)

	if err != nil {
		fmt.Println("Error starting client:", err.Error())

		return
	}

	defer client.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Enter request:")
		request, err := reader.ReadBytes('\n')

		if err != nil {
			fmt.Printf("Error reading request: %v\n", err)
			continue
		}

		resp, err := client.Send(request)

		if len(resp) == 0 {
			continue
		}

		if err != nil {
			fmt.Printf("unable to send request: %v\n", err)
			continue
		}

		fmt.Println("Server response: ", string(resp[:]))
	}
}
