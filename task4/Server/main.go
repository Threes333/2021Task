package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln("13", err)
	}
	for {
		conn, err := l.Accept()
		go func() {
			if err != nil {
				log.Println("19", err)
				return
			}
			rec := make([]byte, 1024)
			_, err = conn.Read(rec)
			if err != nil {
				log.Println("24", err)
				return
			}
			fmt.Println("客户端发出的信息：", string(rec))
			_, err = conn.Write([]byte("hello"))
			if err != nil {
				log.Println("29", err)
				return
			}
		}()
	}
}
