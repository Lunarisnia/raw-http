package main

import (
	"bytes"
	"io"
	"net"
	"os"
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
			// fmt.Println("Hei there")
			var buf bytes.Buffer
			buf.WriteString("Hello, Woerld")
			limited := io.LimitReader(&buf, 3)

			io.Copy(os.Stdout, limited)
			if err != nil {
				panic(err)
			}
			defer c.Close()
			io.Copy(c, c)
		}(conn)
	}
}
