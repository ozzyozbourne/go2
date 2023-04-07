package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"runtime"
	"strings"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	if lis, err := net.Listen("tcp", "localhost:8080"); err != nil {
		log.Fatalln(err)
	} else {
		for {
			fmt.Printf("\n%s\n\n", strings.Repeat(" -*- ", 15))
			fmt.Printf("GOARCH: %s\nGOOS: %s\nMAXPROCS: %d\nNO_OF_CSPS: %d\nN0_OF_CORE_FOR_SCHEDULING: %d\n",
				runtime.GOARCH, runtime.GOOS, runtime.NumCPU(), runtime.NumGoroutine(), runtime.GOMAXPROCS(0))
			if conn, err := lis.Accept(); err != nil {
				log.Printf("Error in accepting a request ")
			} else {
				go handleConn(conn)
			}
		}
	}
}

func handleConn(c net.Conn) {
	defer func(c net.Conn) {
		err := c.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(c)
	for {
		if _, err := io.WriteString(c, "Hello From the server\n"); err != nil {
			log.Fatalln(err)
		} else {
			//	log.Printf("The number of bytes written -> %d\n", i)
		}
		time.Sleep(time.Second)
	}
}
