package main

import (
	"bufio"
	"net"
)

func DoServer(address string) {
	for {
		listen, err := net.Listen("tcp", address)
		if err != nil {
			panic(err)
		}
		conn, err := listen.Accept()
		if err != nil {
			panic(err)
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			panic(err)
		}
		b := []byte(readString)
		conn.Write(b)
	}
}
