package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)
	go setList(channel)

	for v := range channel {
		fmt.Println("recebendo: ", v)
		time.Sleep(time.Second)
	}
}

func setList(channel chan<- int) {
	for i := 0; i < 100; i++ {
		channel <- i
		fmt.Println("enviando: ", i)
	}
	close(channel)
}
