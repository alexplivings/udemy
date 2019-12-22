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

	switch method {
	case "GET":
		handleGet(uri, conn)
	case "POST":
		handlePost(uri, conn)
	default:
		// error message
	}

}

func handleGet(uri string, conn net.Conn) {
	var body string
	switch uri {
	case "/":
		body = "Root page, fun times"
	case "/apply":
		body = "Not sure what we are applying but ok sure"
	default:
		body = "Route not supporteddd"
	}

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func handlePost(uri string, conn net.Conn) {
	var body string
	switch uri {
	case "/apply":
		body = "Not sure what we are applying but ok sure,<br> WITH THE POST METHOD OMG WOW"
	default:
		body = "Route not supporteddd"
	}

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
