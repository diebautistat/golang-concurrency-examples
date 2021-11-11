package main

import (
	"fmt"
)

func counter(limit int, c chan int) {
	for i := 0; i < limit; i++ {
		c <- i
	}
	close(c)
}

func main() {
	channel := make(chan int, 10)
	go counter(10, channel)
	for i := range channel {
		fmt.Println(i)
	}
}

