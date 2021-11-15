package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	m, u := request(conn)
	respond(conn, m, u)
}

func request(conn net.Conn) (m, u string) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		method := strings.Fields(ln)[0]
		route := strings.Fields(ln)[1]
		return method, route
	}
	return "", ""
}

func respond(conn net.Conn, method string, url string) {
	var dsp string
	switch method {
	case "GET":
		dsp = "Deci vrei sa primesti date?"
	case "POST":
		dsp = "Deci vrei sa inserezi date?"
	}
	dsp = dsp + " si ai ajuns la url acesta : " + url
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>` + dsp + `</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

}
