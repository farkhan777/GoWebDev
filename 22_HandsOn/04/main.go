package main

import (
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	io.WriteString(conn, "I see you connected")
	defer conn.Close()
}
