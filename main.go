package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func(c net.Conn) {
			if err != nil {
				panic(err)
			}
			defer c.Close()
			httpStatus := "HTTP/1.1 200 OK\r\n"
			content := "Hey there"
			httpContent := fmt.Sprintf(
				"%sContent-Length: %v\r\n\r\n%s", httpStatus,
				len(content), content,
			)

			fmt.Fprintf(c, httpContent)
		}(conn)
	}
}
