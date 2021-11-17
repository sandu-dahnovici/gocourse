package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
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
		defer conn.Close()
		fmt.Fprint(conn, "I see you connected")
		r := bufio.NewScanner(conn)
		for r.Scan() {
			ln := r.Text()
			io.WriteString(os.Stdout, ln)
		}
		break
	}
}
