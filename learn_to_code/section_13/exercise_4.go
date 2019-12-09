package main

import "fmt"

func main() {
	config := struct {
		retries    int
		serverName string
	}{
		retries:    10,
		serverName: "test_server.test",
	}

	fmt.Println(config)
}
