package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		log.Fatalln("13", err)
	}
	_, err = conn.Write([]byte("hi"))
	if err != nil {
		log.Fatalln("17", err)
	}
	for {
		msg := make([]byte, 1024)
		_, err = conn.Read(msg)
		fmt.Println("服务端发出的信息：", string(msg))
		if err != nil {
			log.Println(err)
		} else {
			break
		}
	}
}
