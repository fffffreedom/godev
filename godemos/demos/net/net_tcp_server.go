package main

import (
	"fmt"
	"net"
)

func doServerStuff(conn net.Conn) {
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			return //终止程序
		}
		fmt.Printf("Received data: %v\n", string(buf[:len]))
	}
}

func main() {
	fmt.Println("starting the server...")

	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("tcp listen error: ", err)
		return
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("tcp Accept error!")
			return
		}
		go doServerStuff(conn)
	}
}
