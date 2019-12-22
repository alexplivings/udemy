package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	var method, uri string

	i := 0
	for scanner.Scan() {
		ln := scanner.Text()
		if ln == "" {
			break
		}
		if i == 0 {
			str := strings.Fields(ln)
			method = str[0]
			uri = str[1]
		}
		fmt.Println(ln)
		i++
	}

	body := fmt.Sprintf("Method: %v<br>URI: %v", method, uri)

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)

}
