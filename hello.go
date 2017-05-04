package main

import (
	"net"
	"log"
	"fmt"
	"time"
)


func f() *int {
	a := 10
	return &a
}

func main() {
	conn, err := net.DialTimeout("tcp", "192.168.150.63:8089", 5 * time.Second)
	if err != nil {
		log.Fatalf("Failed to connection, err:%v", err)
	}

	i := 0
	for {
		byteLen, err := conn.Write([]byte("hello world"))
		if err != nil {
			log.Printf("Server close connection[srv: %s]", conn.RemoteAddr().String())
			conn.Close()
			return
		}
		fmt.Println(byteLen)
		time.Sleep(2 * time.Second)
		i++
		if 10 == i {
			break
		}
	}
	conn.Close()
}