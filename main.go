package main

import (
	"fmt"
	"github.com/efegurkan/dummyhttp/api"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Starting application\r\n")
	givenport := os.Args[1:]

	if len(givenport) > 0 {
		_, err := strconv.Atoi(givenport[0])
		if err != nil {
			fmt.Printf("Cannot parse port number \r\n")
		} else {
			api.InitializeServer(givenport[0])
		}
	} else {
		api.InitializeServer("1234")
	}
}
