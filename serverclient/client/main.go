package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			log.Printf("Error in close the connection at the client side\n")
			return
		}
	}(conn)
	mustCopy(os.Stdout, conn)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if i, err := io.Copy(dst, src); err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("The Number of bytes Writen from source to destination in client -> %d", i)
	}
}
