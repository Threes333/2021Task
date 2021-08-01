package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Producer(out chan<- int) {
	for {
		out <- rand.Int()
		time.Sleep(time.Second)
	}
}

func Consumer(in <-chan int) {
	for {
		select {
		case v := <-in:
			fmt.Println(v)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	channel := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go Producer(channel)
	}
	Consumer(channel)
}
